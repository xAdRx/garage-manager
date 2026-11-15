import type { VehicleFilters, VehicleType } from '../api/vehicles.ts'

interface Props {
  filters: VehicleFilters
  onChange: (filters: VehicleFilters) => void
  onReset: () => void
}

const typeOptions: { value: VehicleType | ''; label: string }[] = [
  { value: '', label: 'All' },
  { value: 'car', label: 'Cars' },
  { value: 'motorcycle', label: 'Motorcycles' },
]

export default function FilterBar({ filters, onChange, onReset }: Props) {
  const set = (key: keyof VehicleFilters, value: string) => onChange({ ...filters, [key]: value })
  const hasFilters = Object.values(filters).some((v) => v.trim() !== '')

  return (
    <section className="filters" aria-label="Filter vehicles">
      <div className="filters__group">
        <span className="field__label" id="filter-type-label">
          Type
        </span>
        <div className="segmented" role="radiogroup" aria-labelledby="filter-type-label">
          {typeOptions.map((opt) => (
            <button
              key={opt.value || 'all'}
              type="button"
              role="radio"
              aria-checked={filters.type === opt.value}
              className="segmented__option"
              onClick={() => set('type', opt.value)}
            >
              {opt.label}
            </button>
          ))}
        </div>
      </div>

      <label className="field">
        <span className="field__label">Brand</span>
        <input
          className="input"
          type="search"
          placeholder="e.g. chev"
          value={filters.brand}
          onChange={(e) => set('brand', e.target.value)}
        />
      </label>

      <label className="field">
        <span className="field__label">Model</span>
        <input
          className="input"
          type="search"
          placeholder="e.g. corvette"
          value={filters.model}
          onChange={(e) => set('model', e.target.value)}
        />
      </label>

      <div className="field field--range">
        <span className="field__label" id="year-range-label">
          Year
        </span>
        <div className="range" role="group" aria-labelledby="year-range-label">
          <input
            className="input input--year"
            type="number"
            inputMode="numeric"
            placeholder="From"
            aria-label="Year from"
            value={filters.minYear}
            onChange={(e) => set('minYear', e.target.value)}
          />
          <span className="range__dash" aria-hidden="true">
            –
          </span>
          <input
            className="input input--year"
            type="number"
            inputMode="numeric"
            placeholder="To"
            aria-label="Year to"
            value={filters.maxYear}
            onChange={(e) => set('maxYear', e.target.value)}
          />
        </div>
      </div>

      <button type="button" className="button button--ghost filters__reset" onClick={onReset} disabled={!hasFilters}>
        Clear filters
      </button>
    </section>
  )
}
