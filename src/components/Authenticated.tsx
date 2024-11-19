import { Suspense, useContext, useEffect } from "react";
import { Outlet, useNavigate } from "react-router-dom";

import { Loading } from ".";
import { AuthContext } from "./Context";

export default function User() {
  const { me, meLoading, meError } = useContext(AuthContext);
  const navigate = useNavigate();

  const urlParameters = new URLSearchParams(location.search);

  useEffect(() => {
    if (meError && location.pathname !== "/login") {
      urlParameters.set("next", location.pathname);
      navigate(`/login?${urlParameters.toString()}`);
    }
  }, [meError, navigate]);

  if (!me || meLoading) {
    return <Loading />;
  }

  return (
    <Suspense fallback={<Loading />}>
      <Outlet />
    </Suspense>
  );
}
