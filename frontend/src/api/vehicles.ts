// Client for the Go REST API. In development, Vite forwards /api/* to the
// Go server (see vite.config.ts).
const API_BASE = '/api'

export type VehicleType = 'car' | 'motorcycle'

export interface Vehicle {
  id: string
  type: VehicleType
  brand: string
  model: string
  year: number
  doors?: number // cars only, may be missing
  hasSidecar?: boolean // motorcycles only
}

// Filter values as typed in the inputs. Empty string means "no filter".
export interface VehicleFilters {
  type: VehicleType | ''
  brand: string
  model: string
  minYear: string
  maxYear: string
}

export const emptyFilters: VehicleFilters = {
  type: '',
  brand: '',
  model: '',
  minYear: '',
  maxYear: '',
}

export interface NewVehicle {
  type: VehicleType
  brand: string
  model: string
  year: number
  doors?: number
  hasSidecar?: boolean
}

// ApiError carries the API's plain-text error message and HTTP status.
export class ApiError extends Error {
  readonly status: number

  constructor(message: string, status: number) {
    super(message)
    this.name = 'ApiError'
    this.status = status
  }
}

async function parseResponse<T>(res: Response): Promise<T> {
  if (!res.ok) {
    const text = (await res.text()).trim()
    throw new ApiError(text || `Request failed with status ${res.status}`, res.status)
  }
  return (await res.json()) as T
}

export async function fetchVehicles(filters: VehicleFilters, signal?: AbortSignal): Promise<Vehicle[]> {
  const params = new URLSearchParams()
  for (const [key, value] of Object.entries(filters)) {
    const trimmed = value.trim()
    if (trimmed !== '') params.set(key, trimmed)
  }
  const query = params.size > 0 ? `?${params}` : ''
  const res = await fetch(`${API_BASE}/vehicles${query}`, { signal })
  return parseResponse<Vehicle[]>(res)
}

export async function createVehicle(vehicle: NewVehicle): Promise<Vehicle> {
  const res = await fetch(`${API_BASE}/vehicles`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(vehicle),
  })
  return parseResponse<Vehicle>(res)
}
