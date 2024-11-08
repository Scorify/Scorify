import { Box, Container, Typography } from "@mui/material";

export default function Statuses() {
  return (
    <Container component='main' maxWidth='md'>
      <Box
        sx={{
          marginTop: 8,
          display: "flex",
          flexDirection: "column",
          alignItems: "center",
        }}
      >
        <Typography component='h1' variant='h3' sx={{ marginBottom: "24px" }}>
          Status Query
        </Typography>
      </Box>
    </Container>
  );
}
