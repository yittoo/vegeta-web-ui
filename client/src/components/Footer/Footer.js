import React from "react";
import { Row, Col } from "antd";

import s from "./Footer.module.scss";

export class Footer extends React.Component {
  render() {
    return (
      <footer>
        <Row align="middle" justify="center">
          <Col>Built by Yiğit Sözer</Col>
        </Row>
      </footer>
    );
  }
}
