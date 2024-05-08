import { useAuth } from "react-oidc-context";
import { useEffect, useState } from "react";

export function Home() {
  const auth = useAuth();
  const [res, setRes] = useState(null);

  useEffect(() => {
    (async () => {
      try {
        const token = auth.user?.access_token;
        if (!token) {
          return;
        }
        const response = await fetch(
          "http://localhost:8080/authenticated-ping",
          {
            headers: {
              Authorization: `${token}`,
            },
          }
        );
        setRes(await response.json());
      } catch (e) {
        console.error(e);
      }
    })();
  }, [auth]);

  switch (auth.activeNavigator) {
    case "signinSilent":
      return <div>Signing you in...</div>;
    case "signoutRedirect":
      return <div>Signing you out...</div>;
  }

  if (auth.isLoading) {
    return <div>Loading...</div>;
  }

  if (auth.error) {
    return <div>Oops... {auth.error.message}</div>;
  }

  if (auth.isAuthenticated) {
    return (
      <div
        style={{
          display: "flex",
          flexDirection: "column",
          gap: "32px",
        }}
      >
        Hello {auth.user?.profile.sub}{" "}
        <button onClick={() => void auth.removeUser()}>Log out</button>
        <div>{auth?.user.id_token}</div>
        <div>Server ID: {res?.claims.sub}</div>
      </div>
    );
  }

  return <button onClick={() => void auth.signinRedirect()}>Log in</button>;
}
