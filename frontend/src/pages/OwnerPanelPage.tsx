import { useState, type FormEvent } from 'react'
import { Link } from 'react-router-dom'
import { createVehicle, type NewVehicle, type Vehicle, type VehicleType } from '../api/vehicles.ts'

// Same limits as the API (handlers/vehicle_request.go).
const MIN_YEAR = 1886
const MAX_YEAR = new Date().getFullYear() + 1

interface FormValues {
  type: VehicleType
  brand: string
  model: string
  year: string
  doors: string
  hasSidecar: boolean
}

type FormErrors = Partial<Record<'brand' | 'model' | 'year' | 'doors', string>>

const initialValues: FormValues = {
  type: 'car',
  brand: '',
  model: '',
  year: '',
  doors: '',
  hasSidecar: false,
}

function isWholeNumber(value: string): boolean {
  return /^\d+$/.test(value.trim())
}

function validate(values: FormValues): FormErrors {
  const errors: FormErrors = {}
  if (values.brand.trim() === '') errors.brand = 'Brand is required.'
  if (values.model.trim() === '') errors.model = 'Model is required.'

  if (values.year.trim() === '') {
    errors.year = 'Year is required.'
  } else if (!isWholeNumber(values.year) || +values.year < MIN_YEAR || +values.year > MAX_YEAR) {
    errors.year = `Year must be between ${MIN_YEAR} and ${MAX_YEAR}.`
  }

  if (values.type === 'car' && values.doors.trim() !== '') {
    if (!isWholeNumber(values.doors) || +values.doors < 1 || +values.doors > 10) {
      errors.doors = 'Doors must be between 1 and 10.'
    }
  }
  return errors
}

function toRequest(values: FormValues): NewVehicle {
  const request: NewVehicle = {
    type: values.type,
    brand: values.brand.trim(),
    model: values.model.trim(),
    year: Number(values.year),
  }
  if (values.type === 'car' && values.doors.trim() !== '') request.doors = Number(values.doors)
  if (values.type === 'motorcycle') request.hasSidecar = values.hasSidecar
  return request
}

export default function OwnerPanelPage() {
  const [values, setValues] = useState<FormValues>(initialValues)
  const [errors, setErrors] = useState<FormErrors>({})
  const [submitting, setSubmitting] = useState(false)
  const [serverError, setServerError] = useState<string | null>(null)
  const [created, setCreated] = useState<Vehicle | null>(null)

  const set = <K extends keyof FormValues>(key: K, value: FormValues[K]) => {
    setValues((prev) => ({ ...prev, [key]: value }))
    if (key in errors) setErrors((prev) => ({ ...prev, [key]: undefined }))
  }

  async function handleSubmit(e: FormEvent<HTMLFormElement>) {
    e.preventDefault()
    const found = validate(values)
    setErrors(found)
    setServerError(null)
    setCreated(null)
    if (Object.keys(found).length > 0) return

    setSubmitting(true)
    try {
      const vehicle = await createVehicle(toRequest(values))
      setCreated(vehicle)
      // Keep the chosen type so several vehicles of one kind are quick to add.
      setValues({ ...initialValues, type: values.type })
    } catch (err) {
      setServerError(err instanceof Error ? err.message : 'Unknown error')
    } finally {
      setSubmitting(false)
    }
  }

  return (
    <>
      <div className="page-title">
        <h1>Owner Panel</h1>
        <p className="page-title__count">Register a new vehicle in the garage</p>
      </div>

      <section className="panel" aria-labelledby="add-vehicle-title">
        <div className="panel__header">
          <h2 id="add-vehicle-title">Add vehicle</h2>
          <span className="panel__tag">Owner only</span>
        </div>

        <div aria-live="polite">
          {created && (
            <div className="notice notice--success">
              <p className="notice__title">
                {created.brand} {created.model} ({created.year}) is parked in the garage.
              </p>
              <Link to="/" className="notice__link">
                View it in the garage →
              </Link>
            </div>
          )}
          {serverError && (
            <div className="notice notice--error" role="alert">
              <p className="notice__title">The API rejected the vehicle</p>
              <p>{serverError}</p>
            </div>
          )}
        </div>

        <form className="form" onSubmit={handleSubmit} noValidate>
          <div className="form__row">
            <span className="field__label" id="form-type-label">
              Type
            </span>
            <div className="segmented" role="radiogroup" aria-labelledby="form-type-label">
              {(['car', 'motorcycle'] as const).map((t) => (
                <button
                  key={t}
                  type="button"
                  role="radio"
                  aria-checked={values.type === t}
                  className="segmented__option"
                  onClick={() => set('type', t)}
                >
                  {t === 'car' ? 'Car' : 'Motorcycle'}
                </button>
              ))}
            </div>
          </div>

          <div className="form__grid">
            <TextField
              label="Brand"
              required
              value={values.brand}
              error={errors.brand}
              placeholder="e.g. Skoda"
              onChange={(v) => set('brand', v)}
            />
            <TextField
              label="Model"
              required
              value={values.model}
              error={errors.model}
              placeholder="e.g. Octavia"
              onChange={(v) => set('model', v)}
            />
            <TextField
              label="Year"
              required
              numeric
              value={values.year}
              error={errors.year}
              placeholder={`${MIN_YEAR}–${MAX_YEAR}`}
              onChange={(v) => set('year', v)}
            />

            {values.type === 'car' ? (
              <TextField
                label="Doors"
                numeric
                value={values.doors}
                error={errors.doors}
                placeholder="1–10"
                onChange={(v) => set('doors', v)}
              />
            ) : (
              <div className="field">
                <span className="field__label">
                  Sidecar <span className="field__optional">optional</span>
                </span>
                <label className="checkbox">
                  <input
                    type="checkbox"
                    checked={values.hasSidecar}
                    onChange={(e) => set('hasSidecar', e.target.checked)}
                  />
                  <span>Has a sidecar</span>
                </label>
              </div>
            )}
          </div>

          <div className="form__actions">
            <button type="submit" className="button button--primary" disabled={submitting}>
              {submitting ? 'Parking…' : 'Add to garage'}
            </button>
            <button
              type="button"
              className="button button--ghost"
              disabled={submitting}
              onClick={() => {
                setValues(initialValues)
                setErrors({})
                setServerError(null)
              }}
            >
              Clear form
            </button>
          </div>
        </form>
      </section>
    </>
  )
}

interface TextFieldProps {
  label: string
  value: string
  onChange: (value: string) => void
  error?: string
  placeholder?: string
  required?: boolean
  numeric?: boolean
}

function TextField({ label, value, onChange, error, placeholder, required, numeric }: TextFieldProps) {
  const id = `field-${label.toLowerCase()}`
  return (
    <div className="field">
      <label className="field__label" htmlFor={id}>
        {label} {required ? <span className="field__required" aria-hidden="true">*</span> : <span className="field__optional">optional</span>}
      </label>
      <input
        id={id}
        className={`input${error ? ' input--invalid' : ''}`}
        type="text"
        inputMode={numeric ? 'numeric' : undefined}
        value={value}
        placeholder={placeholder}
        required={required}
        aria-invalid={error ? true : undefined}
        aria-describedby={error ? `${id}-error` : undefined}
        onChange={(e) => onChange(e.target.value)}
      />
      {error && (
        <p id={`${id}-error`} className="field__error">
          {error}
        </p>
      )}
    </div>
  )
}
