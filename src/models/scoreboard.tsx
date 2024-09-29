import { Property } from "csstype";

import { ScoreboardQuery, StatusEnum } from "../graph";

export type ScoreboardData = {
  top: (number | null | undefined)[];
  left: (string | null | undefined)[];
  values: ScoreboardQuery["scoreboard"]["statuses"];
};

export type ScoreboardTheme = {
  heading: {
    dark: {
      highlighted: Property.BackgroundColor;
      plain: Property.BackgroundColor;
    };
    light: {
      highlighted: Property.BackgroundColor;
      plain: Property.BackgroundColor;
    };
  };
  cell: {
    dark: {
      highlighted: {
        [StatusEnum.Down]: Property.BackgroundColor;
        [StatusEnum.Up]: Property.BackgroundColor;
        [StatusEnum.Unknown]: Property.BackgroundColor;
      };
      plain: {
        [StatusEnum.Down]: Property.BackgroundColor;
        [StatusEnum.Up]: Property.BackgroundColor;
        [StatusEnum.Unknown]: Property.BackgroundColor;
      };
    };
    light: {
      highlighted: {
        [StatusEnum.Down]: Property.BackgroundColor;
        [StatusEnum.Up]: Property.BackgroundColor;
        [StatusEnum.Unknown]: Property.BackgroundColor;
      };
      plain: {
        [StatusEnum.Down]: Property.BackgroundColor;
        [StatusEnum.Up]: Property.BackgroundColor;
        [StatusEnum.Unknown]: Property.BackgroundColor;
      };
    };
  };
};
