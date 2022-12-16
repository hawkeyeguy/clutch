import * as React from "react";
import { MemoryRouter } from "react-router";
import type { Meta } from "@storybook/react";

import { ApplicationContext } from "../../Contexts/app-context";
import { Grid } from "../../Layout";
import Drawer from "../drawer";
import Header from "../header";

export default {
  title: "Core/AppLayout/Drawer",
  component: Drawer,
  decorators: [
    StoryFn => (
      <MemoryRouter>
        <StoryFn />
      </MemoryRouter>
    ),
    StoryFn => {
      return (
        <ApplicationContext.Provider
          // eslint-disable-next-line react/jsx-no-constructed-context-values
          value={{
            workflows: [
              {
                developer: { name: "Lyft", contactUrl: "mailto:hello@clutch.sh" },
                displayName: "EC2",
                group: "AWS",
                path: "ec2",
                routes: [
                  {
                    component: () => null,
                    componentProps: { resolverType: "clutch.aws.ec2.v1.Instance" },
                    description: "Terminate an EC2 instance.",
                    displayName: "Terminate Instance",
                    path: "instance/terminate",
                    requiredConfigProps: ["resolverType"],
                    trending: true,
                  },
                  {
                    component: () => null,
                    componentProps: { resolverType: "clutch.aws.ec2.v1.AutoscalingGroup" },
                    description: "Resize an autoscaling group.",
                    displayName: "Resize Autoscaling Group",
                    path: "asg/resize",
                    requiredConfigProps: ["resolverType"],
                  },
                ],
              },
            ],
          }}
        >
          <StoryFn />
        </ApplicationContext.Provider>
      );
    },
  ],
  parameters: {
    layout: "fullscreen",
  },
} as Meta;

export const Primary = () => <Drawer />;

export const WithHeader = () => (
  <Grid container direction="column">
    <Header />
    <Drawer />
  </Grid>
);
