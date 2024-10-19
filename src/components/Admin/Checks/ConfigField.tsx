import { Checkbox, FormControlLabel, TextField } from "@mui/material";
import { SxProps, Theme } from "@mui/system";
import { SchemaFieldType } from "../../../graph";

type props = {
  handleInputChange: (key: string, value: number | string | boolean) => void;
  fieldName: string;
  fieldType: SchemaFieldType;
  checkConfig: { [key: string]: string | number | boolean };
  defaultValue?: string;
  enumValues?: string[];
  sx?: SxProps<Theme>;
};

const strToBool = (value: string) => value.toLowerCase() == "true";
const strToNumber = (value: string) => parseInt(value);

export default function ConfigField({
  handleInputChange,
  fieldName,
  fieldType,
  checkConfig,
  defaultValue,
  enumValues,
  sx,
}: props) {
  if (fieldType === SchemaFieldType.Bool) {
    return (
      <FormControlLabel
        label={fieldName}
        control={
          <Checkbox
            checked={
              checkConfig[fieldName] != undefined
                ? !!checkConfig[fieldName]
                : defaultValue !== undefined
                ? strToBool(defaultValue)
                : false
            }
            onChange={(e) =>
              handleInputChange(fieldName, e.target.checked as boolean)
            }
            sx={sx}
          />
        }
      />
    );
  } else if (fieldType === SchemaFieldType.Int) {
    return (
      <TextField
        label={fieldName}
        type='number'
        value={
          checkConfig[fieldName] ??
          (defaultValue !== undefined ? parseInt(defaultValue) : 0)
        }
        onChange={(e) => handleInputChange(fieldName, parseInt(e.target.value))}
        variant='outlined'
        margin='normal'
        sx={sx}
      />
    );
  } else {
    return (
      <TextField
        label={fieldName}
        type='text'
        multiline
        value={defaultValue || checkConfig[fieldName] || ""}
        onChange={(e) => handleInputChange(fieldName, e.target.value)}
        variant='outlined'
        margin='normal'
        sx={sx}
      />
    );
  }
}
