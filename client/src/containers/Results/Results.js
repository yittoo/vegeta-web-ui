import React from "react";
import { Row, Col, Tabs } from "antd";

import s from "./Results.module.scss";
import { ViewResultsConsumer } from "../../context";

const { TabPane } = Tabs;

export class Results extends React.Component {
  renderTabs = (data) => {
    if (!data || data.length === 0) {
      return (
        <Tabs>
          <TabPane key="no_attack" tab="Attack name">
            Results will go here
          </TabPane>
        </Tabs>
      );
    }

    return (
      <Tabs>
        {data.map((a) => (
          <TabPane key={a.timeStampMilisecond} tab={a.name}>
            {this.renderResAsHtml(a.html)}
          </TabPane>
        ))}
      </Tabs>
    );
  };

  renderResAsHtml = (html) => {
    return <iframe className={s.Graph} srcDoc={html} />;
  };

  renderResAsTable = (data) => {
    if (!data || data.length === 0) {
      return <div>Results will go here</div>;
    }
    return (
      <div>
        <h3>Tabular Results</h3>
        {data.map((a) => {
          console.log(a.json);
          return <Row>Soem value</Row>;
        })}
      </div>
    );
  };

  render() {
    const { data } = this.props;
    return (
      <section className={s.Results}>
        <ViewResultsConsumer>
          {({ state }) => {
            if (state.value === state.options.graph) {
              return this.renderTabs(data);
            } else {
              return this.renderResAsTable(data);
            }
          }}
        </ViewResultsConsumer>
      </section>
    );
  }
}
