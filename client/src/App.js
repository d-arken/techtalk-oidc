import logo from "./logo.svg";
import "./App.css";
import { AuthProvider } from "react-oidc-context";
import { Home } from "./Home";

const oidcConfig = {
  authority: "https://darkmdev.kinde.com",
  client_id: "d8db0051dc7647c49da254d35c736f37",
  redirect_uri: "http://localhost:3000",
};

function App() {
  return (
    <AuthProvider {...oidcConfig}>
      <div className="App">
        <Home />
      </div>
    </AuthProvider>
  );
}

export default App;
