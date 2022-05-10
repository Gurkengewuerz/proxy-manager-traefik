import {useSearchParams} from "react-router-dom";
import {useEffect} from "react";
import {getAuthInfo} from "../service/AuthInfo";

const Comp = ({setAuthData}) => {
  const [searchParams, setSearchParams] = useSearchParams();
  const token = searchParams.get("token");

  useEffect(() => {
    console.log(token);
    if (token) {
      localStorage.setItem("authenticationToken", token)
      window.location = "/";
    } else getAuthInfo().then(resp => setAuthData(resp.data));
  }, [token]);

  return <></>
}

export default Comp;