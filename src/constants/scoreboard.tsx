import { green, grey, red, yellow } from "@mui/material/colors";
import { StatusEnum } from "../graph";

export const NormalScoreboardTheme = {
  heading: {
    dark: {
      highlighted: grey[800],
      plain: grey[900],
    },
    light: {
      highlighted: grey[300],
      plain: grey[50],
    },
  },
  cell: {
    dark: {
      highlighted: {
        [StatusEnum.Down]: red[400],
        [StatusEnum.Up]: green[400],
        [StatusEnum.Unknown]: yellow[400],
      },
      plain: {
        [StatusEnum.Down]: red[600],
        [StatusEnum.Up]: green[600],
        [StatusEnum.Unknown]: yellow[600],
      },
    },
    light: {
      highlighted: {
        [StatusEnum.Down]: red[600],
        [StatusEnum.Up]: green[600],
        [StatusEnum.Unknown]: yellow[600],
      },
      plain: {
        [StatusEnum.Down]: red[400],
        [StatusEnum.Up]: green[400],
        [StatusEnum.Unknown]: yellow[400],
      },
    },
  },
};
