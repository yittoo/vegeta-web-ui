import React from "react";
import s from "./Main.module.scss";

import { Drawer } from "../../components";
import { VegetaForm } from "../../containers";

export class Main extends React.Component {
  constructor(props) {
    super(props);
    this.state = {};
  }
  render() {
    return <Drawer originalFromLeft={100} leftChildren={<VegetaForm />} />;
  }
}
