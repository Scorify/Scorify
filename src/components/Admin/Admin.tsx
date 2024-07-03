import { Suspense, useContext, useEffect } from "react";
import { Outlet, useNavigate } from "react-router-dom";

import { Error, Loading } from "..";
import { Role } from "../../graph";
import { AuthContext } from "../Context";

export default function Admin() {
  const { me, meLoading, meError } = useContext(AuthContext);
  const navigate = useNavigate();
  const urlParameters = new URLSearchParams(location.search);

  useEffect(() => {
    if (meError && location.pathname !== "/login") {
      urlParameters.set("next", location.pathname);
      navigate(`/login?${urlParameters.toString()}`);
    }
  }, [meError]);

  if (!me || meLoading) {
    return <Loading />;
  }

  if (me.me?.role !== Role.Admin) {
    return <Error code={403} message='Forbidden' />;
  }

  return (
    <Suspense fallback={<Loading />}>
      <Outlet />
    </Suspense>
  );
}
