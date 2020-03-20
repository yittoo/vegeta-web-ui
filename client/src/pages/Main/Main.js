import React from "react";

import { Drawer } from "../../components";

export class Main extends React.Component {
  constructor(props) {
    super(props);
    this.state = {};
  }
  render() {
    return <Drawer originalFromLeft={100} />;
  }
}
