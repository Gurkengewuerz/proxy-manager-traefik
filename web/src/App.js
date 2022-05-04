import {Route, Routes} from "react-router-dom";
import NotFound from "./components/NotFound";

function App() {
  return (
    <div className="page">

      <header className="navbar navbar-expand-md navbar-light d-print-none">
        <div className="container-xl">
          <div className="nav-item dropdown">
            <a href="#" className="nav-link d-flex lh-1 text-reset p-0" data-bs-toggle="dropdown"
               aria-label="Open user menu">
              <span className="avatar avatar-sm" style={{backgroundImage: "/user.png"}}></span>
              <div className="d-none d-xl-block ps-2">
                <div>Paweł Kuna</div>
                <div className="mt-1 small text-muted">UI Designer</div>
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
                  <a className="nav-link" href="/index.html">
                    <span className="nav-link-icon d-md-none d-lg-inline-block">
                      <svg xmlns="http://www.w3.org/2000/svg" className="icon" width="24" height="24"
                           viewBox="0 0 24 24" strokeWidth="2" stroke="currentColor" fill="none" strokeLinecap="round"
                           strokeLinejoin="round"><path stroke="none" d="M0 0h24v24H0z" fill="none"/><polyline
                        points="5 12 3 12 12 3 21 12 19 12"/><path d="M5 12v7a2 2 0 0 0 2 2h10a2 2 0 0 0 2 -2v-7"/><path
                        d="M9 21v-6a2 2 0 0 1 2 -2h2a2 2 0 0 1 2 2v6"/></svg>
                    </span>
                    <span className="nav-link-title">
                      Home
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
                      Help
                    </span>
                  </a>
                  <div className="dropdown-menu">
                    <a className="dropdown-item" href="/index.html">
                      Documentation
                    </a>
                    <a className="dropdown-item" href="/changelog.html">
                      Changelog
                    </a>
                    <a className="dropdown-item" href="https://github.com/tabler/tabler" target="_blank"
                       rel="noreferrer">
                      Source code
                    </a>
                    <a className="dropdown-item text-pink" href="https://github.com/sponsors/codecalm" target="_blank"
                       rel="noreferrer">
                      <svg xmlns="http://www.w3.org/2000/svg" className="icon icon-inline me-1" width="24" height="24"
                           viewBox="0 0 24 24" strokeWidth="2" stroke="currentColor" fill="none" strokeLinecap="round"
                           strokeLinejoin="round">
                        <path stroke="none" d="M0 0h24v24H0z" fill="none"/>
                        <path d="M19.5 13.572l-7.5 7.428l-7.5 -7.428m0 0a5 5 0 1 1 7.5 -6.566a5 5 0 1 1 7.5 6.572"/>
                      </svg>
                      Sponsor project!
                    </a>
                  </div>
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
                <div className="page-pretitle">
                  Overview
                </div>
                <h2 className="page-title">
                  Horizontal layout
                </h2>
              </div>
            </div>
          </div>
        </div>
        <div className="page-body">
          <div className="container-xl">
            <Routes>
              <Route path="*" element={<NotFound/>}/>
            </Routes>
          </div>
        </div>
      </div>

      <footer className="footer footer-transparent d-print-none">
        <div className="container-xl">
          <div className="row text-center align-items-center flex-row-reverse">
            <div className="col-lg-auto ms-lg-auto">
              <ul className="list-inline list-inline-dots mb-0">
                <li className="list-inline-item"><a href="https://github.com/tabler/tabler" target="_blank"
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
