import { useEffect, useState } from 'react'
import { emptyFilters, fetchVehicles, type Vehicle, type VehicleFilters } from '../api/vehicles.ts'
import FilterBar from '../components/FilterBar.tsx'
import VehicleTile from '../components/VehicleTile.tsx'
import { useDebounced } from '../hooks/useDebounced.ts'

// The last finished request. key identifies which filters it was for, so the
// page can tell during render whether a newer request is still loading.
type LoadResult = { key: string; vehicles: Vehicle[] } | { key: string; error: string }

export default function GaragePage() {
  const [filters, setFilters] = useState<VehicleFilters>(emptyFilters)
  const debouncedFilters = useDebounced(filters, 300)
  const [reloadKey, setReloadKey] = useState(0)
  const [result, setResult] = useState<LoadResult | null>(null)

  const requestKey = `${JSON.stringify(debouncedFilters)}#${reloadKey}`

  useEffect(() => {
    const controller = new AbortController()
    fetchVehicles(debouncedFilters, controller.signal)
      .then((vehicles) => setResult({ key: requestKey, vehicles }))
      .catch((err: unknown) => {
        if (controller.signal.aborted) return
        setResult({ key: requestKey, error: err instanceof Error ? err.message : 'Unknown error' })
      })
    return () => controller.abort()
  }, [debouncedFilters, requestKey])

  const loading = result?.key !== requestKey
  const error = !loading && result && 'error' in result ? result.error : null
  // While a new request loads, keep showing the previous tiles (dimmed).
  const vehicles = result && 'vehicles' in result ? result.vehicles : null
  const filtersActive = Object.values(debouncedFilters).some((v) => v.trim() !== '')

  return (
    <>
      <div className="page-title">
        <h1>Garage floor</h1>
        <p className="page-title__count" aria-live="polite">
          {error
            ? ''
            : vehicles === null
              ? 'Loading…'
              : `${vehicles.length} ${vehicles.length === 1 ? 'vehicle' : 'vehicles'}${filtersActive ? ' match' : ' parked'}`}
        </p>
      </div>

      <FilterBar filters={filters} onChange={setFilters} onReset={() => setFilters(emptyFilters)} />

      {error && (
        <div className="notice notice--error" role="alert">
          <p className="notice__title">Can&apos;t load vehicles</p>
          <p>{error}</p>
          <p className="notice__hint">Check that the Go API is running on port 5000.</p>
          <button type="button" className="button button--ghost" onClick={() => setReloadKey((k) => k + 1)}>
            Try again
          </button>
        </div>
      )}

      {vehicles === null && !error && (
        <div className="bays" aria-hidden="true">
          {[0, 1, 2].map((i) => (
            <div key={i} className="bay bay--skeleton" />
          ))}
        </div>
      )}

      {vehicles !== null && !error && vehicles.length === 0 && (
        <div className="notice">
          <p className="notice__title">{filtersActive ? 'No vehicles match these filters' : 'The garage is empty'}</p>
          {filtersActive ? (
            <button type="button" className="button button--ghost" onClick={() => setFilters(emptyFilters)}>
              Clear filters
            </button>
          ) : (
            <p>Add one in the Owner Panel.</p>
          )}
        </div>
      )}

      {vehicles !== null && !error && vehicles.length > 0 && (
        <div className={`bays${loading ? ' bays--stale' : ''}`}>
          {vehicles.map((v, i) => (
            <VehicleTile key={v.id} vehicle={v} bay={i + 1} />
          ))}
        </div>
      )}
    </>
  )
}
