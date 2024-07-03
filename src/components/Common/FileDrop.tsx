import { useDropzone } from "react-dropzone";

import { Typography, Paper, SxProps } from "@mui/material";
import { CloudUpload } from "@mui/icons-material";

type props = {
  onDrop: (files: File[]) => void;
  onError: (error: Error) => void;
  elevation?: number;
  sx?: SxProps;
};

export default function FileDrop({ onDrop, onError, elevation, sx }: props) {
  const { getRootProps, getInputProps, isDragActive } = useDropzone({
    onDrop,
    onError,
  });

  return (
    <Paper
      {...getRootProps()}
      sx={{
        display: "flex",
        justifyContent: "center",
        alignItems: "center",
        minHeight: "75px",
        padding: "12px",
        border: "2px dashed #ccc",
        cursor: "pointer",
        margin: "16px 0px",
        ...sx,
      }}
      elevation={elevation}
    >
      <input {...getInputProps()} />
      {isDragActive ? (
        <Typography variant='h5' sx={{ color: "#999" }}>
          Drop files here...
        </Typography>
      ) : (
        <>
          <CloudUpload
            sx={{
              fontSize: "36px",
              color: "#ccc",
              marginRight: "8px",
            }}
          />
          <Typography variant='h6' sx={{ color: "#999" }}>
            Add Files
          </Typography>
        </>
      )}
    </Paper>
  );
}
