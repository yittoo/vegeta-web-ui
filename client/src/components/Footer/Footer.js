import React from "react";
import { Row, Col } from "antd";

import s from "./Footer.module.scss";

export class Footer extends React.Component {
  render() {
    return (
      <footer className={s.Footer}>
        <Row align="middle" justify="center">
          <Col>
            <h3>
              Built by{" "}
              <a
                href="https://sozerdesign.com"
                target="_blank"
                rel="noopener noreferrer"
                className={s.SourceLink}
              >
                Yiğit Sözer
              </a>{" "}
              -{" "}
              <a
                href="https://github.com/yittoo/vegeta-web-ui"
                target="_blank"
                rel="noopener noreferrer"
                className={s.SourceLink}
              >
                source #
              </a>
            </h3>
            <h5>
              Artwork by{" "}
              <a
                href="https://github.com/egonelbre/gophers"
                target="_blank"
                rel="noopener noreferrer"
                className={s.SourceLink}
              >
                @egonelbre
              </a>
            </h5>
          </Col>
        </Row>
      </footer>
    );
  }
}
