import { useState, FormEvent } from 'react'
import { ComputerDesktopIcon, ArrowPathIcon } from '@heroicons/react/24/solid'
import { useMutateAuth } from '../hooks/useMutateAuth'
import '../Auth.css'

export const Auth = () => {
  const [email, setEmail] = useState('')
  const [pw, setPw] = useState('')
  const [isLogin, setIsLogin] = useState(true)
  const { loginMutation, registerMutation } = useMutateAuth()

  const submitAuthHandler = async (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault()
    if (isLogin) {
      loginMutation.mutate({ email, password: pw })
    } else {
      await registerMutation
        .mutateAsync({ email, password: pw })
        .then(() => loginMutation.mutate({ email, password: pw }))
    }
  }

  return (
    <div className="auth-container">
      <div className="auth-header">
        <ComputerDesktopIcon className="auth-header-icon" />
        <span className="auth-title">my tech blog</span>
        <ComputerDesktopIcon className="auth-header-icon" />
      </div>
      <h2>{isLogin ? 'Login' : 'Create a new account'}</h2>
      <form className="auth-form" onSubmit={submitAuthHandler}>
        <input
          className="auth-input"
          name="email"
          type="email"
          autoFocus
          placeholder="Email address"
          onChange={(e) => setEmail(e.target.value)}
          value={email}
        />
        <input
          className="auth-input"
          name="password"
          type="password"
          placeholder="Password"
          onChange={(e) => setPw(e.target.value)}
          value={pw}
        />
        <button
          className="auth-button"
          disabled={!email || !pw}
          type="submit"
        >
          {isLogin ? 'Login' : 'Sign Up'}
        </button>
      </form>
      <ArrowPathIcon
        onClick={() => setIsLogin(!isLogin)}
        className="auth-switch"
      />
    </div>
  )
}
