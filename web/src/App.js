import {Route, Routes} from "react-router-dom";
import NotFound from "./components/NotFound";
import LoginCallback from "./components/LoginCallback";
import Dashboard, {title as DashboardTitle} from "./components/dashboard/Dashboard";
import Audit, {title as AuditTitle} from "./components/dashboard/Audit";
import Routers, {title as RoutersTitle} from "./components/dashboard/Routers";
import {useState} from "react";

function App() {
  const [authData, setAuthData] = useState({});
  return (
    <div className="page">

      <header className="navbar navbar-expand-md navbar-dark d-print-none">
        <div className="container-xl">
          <button className="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbar-menu">
            <span className="navbar-toggler-icon"></span>
          </button>
          <h1 className="navbar-brand navbar-brand-autodark d-none-navbar-horizontal pe-0 pe-md-3">
            <a href=".">
              <img src="./static/logo.svg" width="110" height="32" alt="traefik Proxy Manager" className="navbar-brand-image" />
            </a>
          </h1>
          <div className="nav-item dropdown">
            <a href="#" className="nav-link d-flex lh-1 text-reset p-0" data-bs-toggle="dropdown"
               aria-label="Open user menu">
              <span className="avatar avatar-sm"></span>
              <div className="d-none d-xl-block ps-2">
                <div>{authData.preferred_username || "User"}</div>
                <div className="mt-1 small text-muted"></div>
              </div>
            </a>
            <div className="dropdown-menu dropdown-menu-end dropdown-menu-arrow">
              <a href="#" className="dropdown-item">Set status</a>
              <a href="#" className="dropdown-item">Profile & account</a>
              <a href="#" className="dropdown-item">Feedback</a>
              <div className="dropdown-divider"></div>
              <a href="#" className="dropdown-item">Settings</a>
              <a href="#" className="dropdown-item">Logout</a>
            </div>
          </div>
        </div>
      </header>

      <div className="navbar-expand-md">
        <div className="collapse navbar-collapse" id="navbar-menu">
          <div className="navbar navbar-light">
            <div className="container-xl">
              <ul className="navbar-nav">
                <li className="nav-item">
                  <a className="nav-link" href="/dashboard">
                    <span className="nav-link-icon d-md-none d-lg-inline-block">
                      <svg xmlns="http://www.w3.org/2000/svg" className="icon" width="24" height="24"
                           viewBox="0 0 24 24" strokeWidth="2" stroke="currentColor" fill="none" strokeLinecap="round"
                           strokeLinejoin="round"><path stroke="none" d="M0 0h24v24H0z" fill="none"/><polyline
                        points="5 12 3 12 12 3 21 12 19 12"/><path d="M5 12v7a2 2 0 0 0 2 2h10a2 2 0 0 0 2 -2v-7"/><path
                        d="M9 21v-6a2 2 0 0 1 2 -2h2a2 2 0 0 1 2 2v6"/></svg>
                    </span>
                    <span className="nav-link-title">
                      Dashboard
                    </span>
                  </a>
                </li>
                <li className="nav-item dropdown">
                  <a className="nav-link dropdown-toggle" href="#navbar-help" data-bs-toggle="dropdown"
                     data-bs-auto-close="outside" role="button" aria-expanded="false">
                    <span className="nav-link-icon d-md-none d-lg-inline-block">
                      <svg xmlns="http://www.w3.org/2000/svg" className="icon" width="24" height="24"
                           viewBox="0 0 24 24" strokeWidth="2" stroke="currentColor" fill="none" strokeLinecap="round"
                           strokeLinejoin="round"><path stroke="none" d="M0 0h24v24H0z" fill="none"/><circle cx="12"
                                                                                                             cy="12"
                                                                                                             r="4"/><circle
                        cx="12" cy="12" r="9"/><line x1="15" y1="15" x2="18.35" y2="18.35"/><line x1="9" y1="15"
                                                                                                  x2="5.65" y2="18.35"/><line
                        x1="5.65" y1="5.65" x2="9" y2="9"/><line x1="18.35" y1="5.65" x2="15" y2="9"/></svg>
                    </span>
                    <span className="nav-link-title">
                      Middlewares
                    </span>
                  </a>
                  <div className="dropdown-menu">
                    <a className="dropdown-item" href="/index.html">
                      Error Provider
                    </a>
                    <a className="dropdown-item" href="/changelog.html">
                      Auth Provider
                    </a>
                    <a className="dropdown-item" href="/changelog.html">
                      Redirect Provider
                    </a>
                    <a className="dropdown-item" href="/changelog.html">
                      Redirect Scheme Provider
                    </a>
                  </div>
                </li>
                <li className="nav-item">
                  <a className="nav-link" href="/dashboard/router">
                    <span className="nav-link-icon d-md-none d-lg-inline-block">
                      <svg xmlns="http://www.w3.org/2000/svg" className="icon" width="44" height="44"
                           viewBox="0 0 24 24" strokeWidth="1.5" stroke="currentColor" fill="none" strokeLinecap="round"
                           strokeLinejoin="round">
                        <path stroke="none" d="M0 0h24v24H0z" fill="none"></path>
                        <path d="M7 18a4.6 4.4 0 0 1 0 -9a5 4.5 0 0 1 11 2h1a3.5 3.5 0 0 1 0 7h-12"></path>
                      </svg>
                    </span>
                    <span className="nav-link-title">
                      Router
                    </span>
                  </a>
                </li>
                <li className="nav-item">
                  <a className="nav-link" href="/dashboard/audit">
                    <span className="nav-link-icon d-md-none d-lg-inline-block">
                      <svg xmlns="http://www.w3.org/2000/svg" className="icon" width="44" height="44"
                           viewBox="0 0 24 24" strokeWidth="1.5" stroke="#2c3e50" fill="none" strokeLinecap="round" strokeLinejoin="round">
                        <path stroke="none" d="M0 0h24v24H0z" fill="none"></path>
                        <path d="M3 19a9 9 0 0 1 9 0a9 9 0 0 1 9 0"></path>
                        <path d="M3 6a9 9 0 0 1 9 0a9 9 0 0 1 9 0"></path>
                        <line x1="3" y1="6" x2="3" y2="19"></line>
                        <line x1="12" y1="6" x2="12" y2="19"></line>
                        <line x1="21" y1="6" x2="21" y2="19"></line>
                      </svg>
                    </span>
                    <span className="nav-link-title">
                      Audit
                    </span>
                  </a>
                </li>
              </ul>
            </div>
          </div>
        </div>
      </div>

      <div className="page-wrapper">
        <div className="container-xl">
          <div className="page-header d-print-none">
            <div className="row g-2 align-items-center">
              <div className="col">
                <h2 className="page-title">
                  <Routes>
                    <Route path="/" element={<>{DashboardTitle}</>}/>
                    <Route path="/dashboard" element={<>{DashboardTitle}</>}/>
                    <Route path="/dashboard/router" element={<>{RoutersTitle}</>}/>
                    <Route path="/dashboard/audit" element={<>{AuditTitle}</>}/>
                  </Routes>
                </h2>
              </div>
            </div>
          </div>
        </div>
        <div className="page-body">
          <div className="container-xl">
            <Routes>
              <Route path="/" element={<Dashboard/>}/>
              <Route path="/dashboard" element={<Dashboard/>}/>
              <Route path="/dashboard/router" element={<Routers/>}/>
              <Route path="/dashboard/audit" element={<Audit/>}/>
              <Route path="*" element={<NotFound/>}/>
            </Routes>
            <LoginCallback setAuthData={setAuthData}/>
          </div>
        </div>
      </div>

      <footer className="footer footer-transparent d-print-none">
        <div className="container-xl">
          <div className="row text-center align-items-center flex-row-reverse">
            <div className="col-lg-auto ms-lg-auto">
              <ul className="list-inline list-inline-dots mb-0">
                <li className="list-inline-item"><a href="https://github.com/Gurkengewuerz/traefik-proxy-manager"
                                                    target="_blank"
                                                    className="link-secondary" rel="noreferrer">Source code</a></li>
              </ul>
            </div>
            <div className="col-12 col-lg-auto mt-3 mt-lg-0">
              <ul className="list-inline list-inline-dots mb-0">
                <li className="list-inline-item">
                  Copyright &copy; 2022
                  <a href="https://mc8051.de" className="link-secondary">mc8051</a>.
                  All rights reserved.
                </li>
                <li className="list-inline-item">
                  Coded with ❤ with <a href="https://tabler.io" className="link-secondary">Tabler</a>
                </li>
              </ul>
            </div>
          </div>
        </div>
      </footer>

    </div>
  );
}

export default App;
