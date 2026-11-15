import { Link, useLocation } from 'react-router-dom'

export default function Header() {
  const onOwnerPanel = useLocation().pathname.startsWith('/owner')

  return (
    <header className="site-header">
      <div className="site-header__inner">
        <Link to="/" className="brand" aria-label="Garage Manager, go to the garage">
          <svg className="brand__icon" viewBox="0 0 32 32" aria-hidden="true">
            <path d="M3 13 16 4l13 9v16H3z" fill="currentColor" />
            <rect x="7" y="14" width="18" height="15" className="brand__door" />
            <g className="brand__slats">
              <rect x="8" y="15.5" width="16" height="2" />
              <rect x="8" y="19" width="16" height="2" />
              <rect x="8" y="22.5" width="16" height="2" />
              <rect x="8" y="26" width="16" height="2" />
            </g>
          </svg>
          <span className="brand__text">
            <span className="brand__name">Garage Manager</span>
            <span className="brand__tagline">Vehicle storage &amp; records</span>
          </span>
        </Link>

        {onOwnerPanel ? (
          <Link to="/" className="button button--ghost">
            <span aria-hidden="true">←</span> Back to garage
          </Link>
        ) : (
          <Link to="/owner" className="button button--primary">
            <svg viewBox="0 0 20 20" className="button__icon" aria-hidden="true">
              <path
                d="M10 2a4 4 0 0 1 4 4v2h1a1 1 0 0 1 1 1v8a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1V9a1 1 0 0 1 1-1h1V6a4 4 0 0 1 4-4Zm0 2a2 2 0 0 0-2 2v2h4V6a2 2 0 0 0-2-2Z"
                fill="currentColor"
              />
            </svg>
            Owner Panel
          </Link>
        )}
      </div>
      <div className="hazard-stripe" aria-hidden="true" />
    </header>
  )
}
