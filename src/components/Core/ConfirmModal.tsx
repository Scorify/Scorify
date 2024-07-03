import { useMemo, useState, ReactElement } from "react";

import { Box, Button, Modal, TextField, Typography } from "@mui/material";

type props = {
  title: string;
  subtitle: string | ReactElement;
  buttonText: string;
  value: string;
  open: boolean;
  setOpen: (open: boolean) => void;
  onConfirm: () => void;
  label?: string;
};

export default function ConfirmModal({
  title,
  subtitle,
  buttonText,
  value,
  open,
  setOpen,
  onConfirm,
  label,
}: props) {
  const [confirm, setConfirm] = useState<string>("");
  const match = useMemo(() => confirm === value, [confirm, value]);

  return (
    <Modal
      open={open}
      onClose={() => {
        setOpen(false);
      }}
    >
      <Box
        sx={{
          position: "absolute",
          top: "25%",
          left: "50%",
          transform: "translate(-50%, -25%)",
          width: "auto",
          maxWidth: "90vw",
          bgcolor: "background.paper",
          border: `1px solid #000`,
          borderRadius: "8px",
          boxShadow: 24,
          p: 4,
        }}
      >
        <Box
          sx={{
            display: "flex",
            flexDirection: "column",
            alignItems: "center",
          }}
        >
          <Typography component='h1' variant='h3'>
            {title}
          </Typography>
          <Typography component='h2' variant='body1'>
            {subtitle}
          </Typography>
          <TextField
            label={label}
            variant='outlined'
            sx={{
              marginTop: "24px",
            }}
            value={confirm}
            onChange={(event) => {
              setConfirm(event.target.value as string);
            }}
          />

          <Button
            variant='contained'
            sx={{ marginTop: "24px" }}
            onClick={onConfirm}
            disabled={!match}
          >
            {buttonText}
          </Button>
        </Box>
      </Box>
    </Modal>
  );
}
