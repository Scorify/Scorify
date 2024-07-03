import { useState, ChangeEventHandler } from "react";

import { Visibility, VisibilityOff } from "@mui/icons-material";
import IconButton from "@mui/material/IconButton";
import InputAdornment from "@mui/material/InputAdornment";
import TextField from "@mui/material/TextField";

type props = {
  value?: string;
  onBlur?: ChangeEventHandler<HTMLInputElement | HTMLTextAreaElement>;
  onChange?: ChangeEventHandler<HTMLInputElement | HTMLTextAreaElement>;
  variant?: "standard" | "filled" | "outlined" | undefined;
  margin?: "none" | "dense" | "normal" | undefined;
  required?: boolean;
  name?: string;
  label?: string;
  id?: string;
  fullWidth?: boolean;
};

export default function PasswordInput({
  onBlur,
  onChange,
  value,
  variant,
  margin,
  required,
  name,
  label,
  id,
  fullWidth,
}: props) {
  const [password, setPassword] = useState<string>(
    value === undefined ? "" : value
  );
  const [prevPassword, setPrevPassword] = useState<string>(
    value === undefined ? "" : value
  );
  const [showPassword, setShowPassword] = useState(false);

  const handleTogglePassword = () => {
    setShowPassword((prevShowPassword) => !prevShowPassword);
  };

  const handleBlur = (
    e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>
  ) => {
    if (onBlur && password !== prevPassword) {
      onBlur(e);
      setPrevPassword(password);
    }
  };

  return (
    <TextField
      type={showPassword ? "text" : "password"}
      variant={variant}
      value={password}
      onChange={(e) => {
        onChange && onChange(e);
        setPassword(e.target.value);
      }}
      onBlur={handleBlur}
      fullWidth={fullWidth}
      margin={margin}
      required={required}
      name={name}
      label={label}
      id={id}
      InputProps={{
        endAdornment: (
          <InputAdornment position='end'>
            <IconButton onClick={handleTogglePassword}>
              {showPassword ? <VisibilityOff /> : <Visibility />}
            </IconButton>
          </InputAdornment>
        ),
      }}
    />
  );
}
