import React from "react";

import s from "./Layout.module.scss";
import { Header, Footer } from "../../components";

export class Layout extends React.Component {
  render() {
    return (
      <div className={s.Layout}>
        <Header />
        {this.props.children}
        <Footer />
      </div>
    );
  }
}
