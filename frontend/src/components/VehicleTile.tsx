import type { Vehicle } from '../api/vehicles.ts'
import VehicleIcon from './VehicleIcon.tsx'

interface Props {
  vehicle: Vehicle
  bay: number // 1-based parking bay number shown on the tile
}

export default function VehicleTile({ vehicle, bay }: Props) {
  const { type, brand, model, year, doors, hasSidecar, id } = vehicle
  const typeLabel = type === 'car' ? 'Car' : 'Motorcycle'

  return (
    <article className={`bay bay--${type}`} aria-label={`${brand} ${model}, ${year}`}>
      <div className="bay__top">
        <span className="bay__number">Bay {String(bay).padStart(2, '0')}</span>
        <span className="bay__type">{typeLabel}</span>
      </div>

      <div className="bay__floor">
        <VehicleIcon type={type} hasSidecar={hasSidecar} />
      </div>

      <div className="bay__info">
        <p className="bay__brand">{brand}</p>
        <h2 className="bay__model">{model}</h2>
        <dl className="bay__specs">
          <div>
            <dt>Year</dt>
            <dd>{year}</dd>
          </div>
          {type === 'car' && (
            <div>
              <dt>Doors</dt>
              <dd>{doors ?? '—'}</dd>
            </div>
          )}
          {type === 'motorcycle' && (
            <div>
              <dt>Sidecar</dt>
              <dd>{hasSidecar ? 'Yes' : 'No'}</dd>
            </div>
          )}
        </dl>
      </div>

      <p className="bay__id" title={id}>
        ID {id.slice(0, 8)}
      </p>
    </article>
  )
}
