import { Checkbox, FormControlLabel, TextField } from "@mui/material";
import { SxProps, Theme } from "@mui/system";
import { SchemaFieldType } from "../../../graph";

type props = {
  handleInputChange: (key: string, value: number | string | boolean) => void;
  index: string;
  value: SchemaFieldType;
  config: { [key: string]: string | number | boolean };
  default?: string | number | boolean;
  sx?: SxProps<Theme>;
};

export default function ConfigField({
  handleInputChange,
  index,
  value,
  config,
  sx,
}: props) {
  if (value === SchemaFieldType.Bool) {
    return (
      <FormControlLabel
        label={index}
        control={
          <Checkbox
            checked={!!config[index]}
            onChange={(e) =>
              handleInputChange(index, e.target.checked as boolean)
            }
            sx={sx}
          />
        }
      />
    );
  } else if (value === SchemaFieldType.Int) {
    return (
      <TextField
        label={index}
        type='number'
        value={config[index] || ""}
        onChange={(e) => handleInputChange(index, parseInt(e.target.value))}
        variant='outlined'
        margin='normal'
        sx={sx}
      />
    );
  } else {
    return (
      <TextField
        label={index}
        type='text'
        multiline
        value={config[index] || ""}
        onChange={(e) => handleInputChange(index, e.target.value)}
        variant='outlined'
        margin='normal'
        sx={sx}
      />
    );
  }
}
