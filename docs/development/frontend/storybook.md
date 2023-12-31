---
title: Storybook
{{ .EditURL }}
---

import useBaseUrl from '@docusaurus/useBaseUrl';
import Image from '@site/src/components/Image';

[Storybook](https://storybook.js.org/) is an open source tool that allows users to build frontend components and pages in isolation.

Leveraging Storybook makes the development and review process for adding or modifying components easier and faster while also providing a convenient place for Clutch developers to reference when building out different frontend features.

If you're interested in learning more about Storybook be sure to read their [introduction guide](https://storybook.js.org/tutorials/intro-to-storybook/).

## Using Clutch's Storybook

<img alt="button story" src={useBaseUrl('img/docs/feature-development/frontend/storybook/button.png ')} variant="centered" />

All available Clutch components can be found on [Clutch's Storybook site](https://storybook.clutch.sh). This site is automatically published on merges to main and should always be up to date.

### Controls

Each component within storybook exposes it's properties that can be controlled (if any) in a controls tab. This allows for developers to input custom values and see how a component will render.

<img alt="button story controls" src={useBaseUrl('img/docs/feature-development/frontend/storybook/controls.png ')} variant="centered" />

### Actions

Components with some form of action (such as `Button`) will emit a log message within the actions tab when the action is invoked. This log message will show both the property name that registered the event notification as well as any parameters passed back by the component.

<img alt="button story actions" src={useBaseUrl('img/docs/feature-development/frontend/storybook/actions.png ')} variant="centered" />

### Documentation

Clutch leverages a [Storybook addon](https://storybook.js.org/addons/@storybook/addon-docs) to automatically generate documentation pages for every component.

<img alt="button story documentation" src={useBaseUrl('img/docs/feature-development/frontend/storybook/documentation.png ')} variant="centered" />

This page is generated from docstrings and props within the source code of the component and therefore should always be up to date.

At the top of the page, under the component name, is an overview of the selected component, which has been generated by parsing the component's docstring.

<img alt="button story documentation overview" src={useBaseUrl('img/docs/feature-development/frontend/storybook/documentation-overview.png ')} variant="centered" />

```tsx title="button.tsx"
/**
 * A button with default themes based on use case.
 */
const Button: React.FC<ButtonProps> = ({ text, variant = "primary", ...props }) => {
```

Following the overview is a rendering of the component using the current prop values, along with a button to show the code that generates this rendering.

<img alt="button story documentation render" src={useBaseUrl('img/docs/feature-development/frontend/storybook/documentation-render.png ')} variant="centered" />

An easy way to use a Clutch component is to enter values for the component's properties that you would like to use and leverage the documentation to generate the source code for your desired component. This snippet of code can be copy and pasted into your own application to produce the same component.

<img alt="button story documentation code snippet" src={useBaseUrl('img/docs/feature-development/frontend/storybook/documentation-code-snippet.png ')} variant="centered" />

Below the rendering is a list of all properites available on the component, each accompanied by a brief description, a default value if any, and the ability to modify the value.

<img alt="button story documentation component props" src={useBaseUrl('img/docs/feature-development/frontend/storybook/documentation-props.png ')} variant="centered" />

Finally, any additional stories for this component are rendered below the props. In Clutch we use these additional stories to demonstrate modifying key pieces of the component. In the Button component you can see that we have stories to highlight the button variants.

<img alt="button story documentation variant stories" src={useBaseUrl('img/docs/feature-development/frontend/storybook/documentation-variant.png ')} variant="centered" />

## Creating Stories

Under the hood Storybook works by aggregating ["stories"](https://storybook.js.org/docs/ember/get-started/whats-a-story/) within a codebase and, depending on their specified location, renders them into the relevant groupings.

### What's a Story?

A story is a short snippet of code that demonstrates how a specific frontend component with the specified properties will be rendered. Each frontend component should have at least one corresponding story but can have as many as needed to demonstrate all the "interesting" states a component can be in. Interesting is intentionally open ended but within Clutch we tend to add a story for the various states or variants a component can be in or have.

### Structure

Clutch's Storybook is currently split by packages. At the moment it only contains components exported by `@clutch-sh/core`.

When adding a Storybook file it should be placed within a `stories` directory that lives alongside the source code of that story and have the format `{component}.stories.tsx`. This format is important as it allows Storybook to discover the stories within that file.

### Writing Stories

All new components should be accompanied by stories. As mentioned earlier this makes reviewing the new component and any future changes easier for the team and enhances the developer experience within Clutch. 

Every story follows the same format:

  1. It has a default export containing a title, component, and any optional properties such as argument types.

    ```tsx
    export default {
      title: "Core/Buttons/Button",
      component: Button,
      argTypes: {
        onClick: { action: "onClick event" },
      },
    } as Meta;
    ```

    In the above example for the button component stories the title is set to `Core/Buttons/Button`, which becomes the path to these stories within the Storybook component tree. `Button` is specified as the component and an optional type for argument types (`argTypes`) is passed to specify the value for the button's `onClick` property. In this case the `onClick` prop is given a special callback type to allow Storybook to route these events to the actions tab.

  2. It has a Template to use for the component.

    ```tsx
    const Template = (props: ButtonProps) => <Button {...props} />;
    ```

    Continuing with the button example, the template above takes in `ButtonProps`, which have been exported from the component file, and renders a `Button` component passing in any specified props.

  3. It has at least a primary story.

    ```tsx
    export const Primary = Template.bind({});
    Primary.args = {
      text: "Continue",
    };
    ```

    The primary story for button is to use all the default props and provide only those that are required. In this example we set the button's text to "Continue".

When writing stories you can run Clutch's Storybook locally and test out changes. Additionally, every pull request with changes to any component file, storybook file, or dependency versions will build a preview of the Storybook site and publish a link as a status on the pull request.

Additional general information on writing stories can be found on [Storybook's website](https://storybook.js.org/docs/ember/writing-stories/introduction).

## Accessibility

Clutch's Storybook leverages the [a11y addon](https://www.npmjs.com/package/@storybook/addon-a11y) to analyze each component and make a determination if it is compliant with accessibility standards. This is particularly helpful when developing components as it provides an easy way to ensure your changes are compliant by running Storybook locally.

<img alt="button story accessibility addon" src={useBaseUrl('img/docs/feature-development/frontend/storybook/a11y-addon.png ')} variant="centered" />
