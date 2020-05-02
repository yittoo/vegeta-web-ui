import React from "react";
import { Row, Col, Tabs } from "antd";

import s from "./Results.module.scss";

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

  render() {
    return (
      <section className={s.Results}>
        {this.renderTabs(this.props.data)}
      </section>
    );
  }
}
