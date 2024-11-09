import { Autocomplete, Chip, SxProps, TextField } from "@mui/material";

type props = {
  label?: string;
  placeholder?: string;
  options: string[];
  selected: string[];
  setSelected: (selected: string[]) => void;
  sx?: SxProps;
};

export default function Multiselect({
  label,
  placeholder,
  options,
  selected,
  setSelected,
  sx,
}: props) {
  return (
    <Autocomplete
      multiple
      options={options}
      value={selected}
      onChange={(_, value) => setSelected(value)}
      sx={sx}
      renderTags={(value, getTagProps) =>
        value.map((tag, index) => (
          <Chip label={tag} {...getTagProps({ index })} />
        ))
      }
      renderInput={(params) => (
        <TextField {...params} label={label} placeholder={placeholder} />
      )}
    />
  );
}
