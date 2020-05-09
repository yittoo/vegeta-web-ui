import React from "react";
import { Row, Col, Switch } from "antd";

import s from "./Header.module.scss";
import Gopher from "../../assets/gopher-square.png";

import { ViewResultsConsumer } from "../../context";

export class Header extends React.Component {
  render() {
    return (
      <header className={s.Header}>
        <Row align="middle" justify="space-between" gutter={20}>
          <Col>
            <Row align="middle" gutter={20}>
              <Col flex="none">
                <img src={Gopher} />
              </Col>
              <Col>
                <h2 className={s.Title}>
                  Vegeta Web UI - Http Load Testing Tool
                </h2>
                <h4 className={s.SubTitle}>
                  Built upon{" "}
                  <a
                    href="https://github.com/tsenart/vegeta"
                    target="_blank"
                    rel="noopener noreferrer"
                    className={s.SourceLink}
                  >
                    work of tsenart
                  </a>
                </h4>
              </Col>
            </Row>
          </Col>
          <Col flex="none">
            <p className={s.OptionsTitle}>View results as</p>
            <Row align="middle" gutter={15}>
              <Col>
                <p className={s.OptionsText}>Table comparison</p>
              </Col>
              <Col>
                <ViewResultsConsumer>
                  {({ state, dispatch, actions }) => {
                    return (
                      <Switch
                        checked={state.value === state.options.graph}
                        onChange={() => dispatch(actions.toggleViewResultsAs())}
                      />
                    );
                  }}
                </ViewResultsConsumer>
              </Col>
              <Col>
                <p className={s.OptionsText}>Individual Graphs</p>
              </Col>
            </Row>
          </Col>
        </Row>
      </header>
    );
  }
}
