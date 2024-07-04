import { Box, Container, Typography } from "@mui/material";

type props = {
  code: number;
  message: string;
};

export default function Error({ code, message }: props) {
  return (
    <Container>
      <Box
        sx={{
          marginTop: 8,
          display: "flex",
          flexDirection: "column",
          alignItems: "center",
        }}
      >
        <Typography component='h1' variant='h3'>
          {message}
        </Typography>
        <Typography component='h1' variant='h1'>
          {code}
        </Typography>
      </Box>
    </Container>
  );
}
