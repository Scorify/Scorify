import { Autocomplete, Chip, TextField } from "@mui/material";

type props = {
  label?: string;
  placeholder?: string;
  options: string[];
  selected: string[];
  setSelected: (selected: string[]) => void;
};

export default function Multiselect({
  label,
  placeholder,
  options,
  selected,
  setSelected,
}: props) {
  return (
    <Autocomplete
      multiple
      options={options}
      value={selected}
      onChange={(_, value) => setSelected(value)}
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
