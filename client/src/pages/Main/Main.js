import React from "react";
import { Row, Col } from "antd";

import s from "./Main.module.scss";
import { VegetaForm, Layout, Results } from "../../containers";

export class Main extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      results: [],
    };
  }

  addNewResult = ({ name, payload }) => {
    let resCopy = [...this.state.results];

    const { json, html, timeStampSecond } = payload;
    resCopy.push({
      name,
      html,
      json,
      timeStampMilisecond: timeStampSecond * 1000,
    });

    this.setState({
      results: resCopy,
    });
  };
  render() {
    return (
      <Layout>
        <Row className={s.Main} gutter={20}>
          <Col span={6}>
            <VegetaForm
              onAddNewResult={({ name, payload }) =>
                this.addNewResult({ name, payload })
              }
            />
          </Col>
          <Col span={18}>
            <Results data={this.state.results} />
          </Col>
        </Row>
      </Layout>
    );
  }
}
