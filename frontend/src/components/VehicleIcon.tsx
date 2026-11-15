import type { VehicleType } from '../api/vehicles.ts'

interface Props {
  type: VehicleType
  hasSidecar?: boolean
}

// Side-view silhouette of a car or motorcycle, drawn in currentColor.
export default function VehicleIcon({ type, hasSidecar }: Props) {
  return (
    <svg className="vehicle-icon" viewBox="0 0 120 60" aria-hidden="true">
      {type === 'car' ? <Car /> : <Motorcycle hasSidecar={hasSidecar} />}
    </svg>
  )
}

function Wheel({ cx, cy, r }: { cx: number; cy: number; r: number }) {
  return (
    <g>
      <circle cx={cx} cy={cy} r={r} className="vehicle-icon__tyre" />
      <circle cx={cx} cy={cy} r={r * 0.35} fill="currentColor" />
    </g>
  )
}

function Car() {
  return (
    <g>
      <path
        d="M8 44V35q1-6 11-7.5l16-1.5 12-11q3-2.5 8-2.5h24q6 0 10 4l9 9.5 9 2q7 1.5 7 8V44Z"
        fill="currentColor"
      />
      <path d="M40 26.5 50.5 17q1.5-1 4-1H66v10.5Z" className="vehicle-icon__glass" />
      <path d="M70 16h10q4 0 6.5 2.5l8 8H70Z" className="vehicle-icon__glass" />
      <rect x="104" y="31" width="7" height="3" rx="1.5" className="vehicle-icon__light" />
      <Wheel cx={30} cy={45} r={9.5} />
      <Wheel cx={92} cy={45} r={9.5} />
    </g>
  )
}

function Motorcycle({ hasSidecar }: { hasSidecar?: boolean }) {
  return (
    <g>
      {hasSidecar && (
        <g className="vehicle-icon__sidecar">
          <path d="M48 42q-2-14 14-15h26q12 0 14 9l1 6Z" fill="currentColor" />
          <circle cx={92} cy={46} r={8} className="vehicle-icon__tyre" />
        </g>
      )}
      <g fill="none" stroke="currentColor" strokeWidth="4" strokeLinecap="round" strokeLinejoin="round">
        <path d="M24 44 44 30h26l22 14" />
        <path d="M92 44 82 16" />
        <path d="M76 15h11" />
      </g>
      <path d="M42 28q4-9 18-8l12 1q4 1 4 7Z" fill="currentColor" />
      <path d="M20 26h22v3H22q-3 0-2-3Z" fill="currentColor" />
      <rect x="50" y="31" width="18" height="11" rx="2" fill="currentColor" />
      <Wheel cx={24} cy={44} r={12} />
      <Wheel cx={92} cy={44} r={12} />
    </g>
  )
}
