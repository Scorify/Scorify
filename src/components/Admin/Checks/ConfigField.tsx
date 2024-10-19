import {
  Checkbox,
  FormControl,
  FormControlLabel,
  InputLabel,
  MenuItem,
  Select,
  TextField,
} from "@mui/material";
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
    if (enumValues) {
      return (
        <FormControl sx={{ width: "214px", mt: "16px", mb: "8px" }}>
          <InputLabel id={fieldName}>{fieldName}</InputLabel>
          <Select
            labelId={fieldName}
            label={fieldName}
            type='number'
            value={
              checkConfig[fieldName] ??
              (defaultValue !== undefined
                ? parseInt(defaultValue)
                : parseInt(enumValues[0]))
            }
            variant='outlined'
            onChange={(e) => handleInputChange(fieldName, e.target.value)}
            sx={sx}
          >
            {enumValues.map((enumValue) => (
              <MenuItem key={enumValue} value={parseInt(enumValue)}>
                {enumValue}
              </MenuItem>
            ))}
          </Select>
        </FormControl>
      );
    }

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
    if (enumValues) {
      return (
        <FormControl sx={{ width: "214px", mt: "16px", mb: "8px" }}>
          <InputLabel id={fieldName}>{fieldName}</InputLabel>
          <Select
            labelId={fieldName}
            label={fieldName}
            type='text'
            value={
              checkConfig[fieldName] ??
              (defaultValue !== undefined ? defaultValue : enumValues[0])
            }
            variant='outlined'
            onChange={(e) => handleInputChange(fieldName, e.target.value)}
            sx={sx}
          >
            {enumValues.map((enumValue) => (
              <MenuItem key={enumValue} value={enumValue}>
                {enumValue}
              </MenuItem>
            ))}
          </Select>
        </FormControl>
      );
    }

    return (
      <TextField
        label={fieldName}
        type='text'
        multiline
        value={
          checkConfig[fieldName] ??
          (defaultValue !== undefined ? defaultValue : "")
        }
        onChange={(e) => handleInputChange(fieldName, e.target.value)}
        variant='outlined'
        margin='normal'
        sx={sx}
      />
    );
  }
}
