import React, { useEffect, useRef } from 'react';
import { routes } from '../constants/routeMap';
import * as d3 from 'd3';

interface Node extends d3.SimulationNodeDatum {
  id: string;
  name: string;
}

interface Link {
  source: string;
  target: string;
}

const RouteMapVisualizer: React.FC = () => {
  const svgRef = useRef<SVGSVGElement>(null);

  useEffect(() => {
    const currentSvgRef = svgRef.current;
    if (!currentSvgRef) return;

    // 既存のSVG内容をクリア
    d3.select(currentSvgRef).selectAll('*').remove();

    // ノードとリンクのデータを準備
    const nodes: Node[] = routes.map(route => ({
      id: route.id,
      name: route.name
    }));

    const links: Link[] = routes.flatMap(route => 
      (route.children || []).map(childId => ({
        source: route.id,
        target: childId
      }))
    );

    // SVGサイズの設定
    const width = 800;
    const height = 600;
    const svg = d3.select(currentSvgRef)
      .attr('width', width)
      .attr('height', height);

    // シミュレーションの設定
    const simulation = d3.forceSimulation<Node>(nodes)
      .force('link', d3.forceLink<Node, Link>(links).id(d => d.id).distance(150))
      .force('charge', d3.forceManyBody().strength(-300))
      .force('center', d3.forceCenter(width / 2, height / 2));

    // 矢印マーカーの定義
    svg.append('defs').append('marker')
      .attr('id', 'arrowhead')
      .attr('viewBox', '0 -5 10 10')
      .attr('refX', 15)
      .attr('refY', 0)
      .attr('orient', 'auto')
      .attr('markerWidth', 6)
      .attr('markerHeight', 6)
      .append('path')
      .attr('d', 'M0,-5L10,0L0,5')
      .attr('fill', '#999');

    // リンクの描画
    const link = svg.append('g')
      .selectAll('line')
      .data(links)
      .enter()
      .append('line')
      .attr('stroke', '#999')
      .attr('stroke-width', 2)
      .attr('marker-end', 'url(#arrowhead)');

    // ノードの描画
    const node = svg.append('g')
      .selectAll('g')
      .data(nodes)
      .enter()
      .append('g')
      .call(d3.drag<SVGGElement, Node>()
        .on('start', dragstarted)
        .on('drag', dragged)
        .on('end', dragended));

    node.append('circle')
      .attr('r', 20)
      .attr('fill', '#69b3a2');

    node.append('text')
      .text(d => d.name)
      .attr('text-anchor', 'middle')
      .attr('dy', 30);

    // シミュレーションの更新関数
    simulation.on('tick', () => {
      link
        .attr('x1', d => (d.source as unknown as Node).x || 0)
        .attr('y1', d => (d.source as unknown as Node).y || 0)
        .attr('x2', d => (d.target as unknown as Node).x || 0)
        .attr('y2', d => (d.target as unknown as Node).y || 0);

      node
        .attr('transform', d => `translate(${d.x || 0},${d.y || 0})`);
    });

    // ドラッグ関連の関数
    function dragstarted(event: any, d: Node) {
      if (!event.active) simulation.alphaTarget(0.3).restart();
      d.fx = d.x;
      d.fy = d.y;
    }

    function dragged(event: any, d: Node) {
      d.fx = event.x;
      d.fy = event.y;
    }

    function dragended(event: any, d: Node) {
      if (!event.active) simulation.alphaTarget(0);
      d.fx = null;
      d.fy = null;
    }

    return () => {
      // クリーンアップ関数を強化
      simulation.stop();
      if (currentSvgRef) {
        d3.select(currentSvgRef).selectAll('*').remove();
      }
    };
  }, []); // 依存配列は空のままで
  return (
    <div className="route-map-container">
      <h2>画面遷移図</h2>
      <svg ref={svgRef}></svg>
    </div>
  );
};

export default RouteMapVisualizer;