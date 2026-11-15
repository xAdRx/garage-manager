import { useEffect, useState } from 'react'

// useDebounced returns value after it has stopped changing for delayMs,
// so typing in a filter does not send a request on every key press.
export function useDebounced<T>(value: T, delayMs: number): T {
  const [debounced, setDebounced] = useState(value)

  useEffect(() => {
    const timer = setTimeout(() => setDebounced(value), delayMs)
    return () => clearTimeout(timer)
  }, [value, delayMs])

  return debounced
}
