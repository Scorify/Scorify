import { Chip } from "@mui/material";

import { File as gqlFile } from "../../graph";

type props = {
  file: gqlFile | File;
  color?:
    | "default"
    | "error"
    | "primary"
    | "secondary"
    | "info"
    | "success"
    | "warning";
  onDelete?: () => void;
  size?: "small" | "medium";
};

export default function FileChip({ file, color, onDelete, size }: props) {
  const label =
    file.name.length > 25
      ? `${file.name.slice(0, 10)}[...]${file.name.slice(
          file.name.length - 10
        )}`
      : file.name;

  const onClick = () => {
    if (file instanceof File) {
      window.open(URL.createObjectURL(file), "_blank");
    } else {
      window.open("http://localhost:8080" + file.url, "_blank");
    }
  };

  return (
    <Chip
      label={label}
      onClick={onClick}
      color={color}
      onDelete={onDelete}
      size={size}
    />
  );
}
