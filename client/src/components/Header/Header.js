import React from "react";

import { Row, Col } from "antd";

export class Header extends React.Component {
  render() {
    return (
      <header>
        <Row align="middle" justify="space-between" gutter={20}>
          <Col>Logo</Col>
          <Col>Options</Col>
        </Row>
      </header>
    );
  }
}
