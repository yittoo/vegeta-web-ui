import React from "react";
import Router from "./router";
import "./App.scss";

import { ViewResultsProvider } from "../context/index";

const App = () => {
  return (
    <ViewResultsProvider>
      <Router />
    </ViewResultsProvider>
  );
};

export default App;
