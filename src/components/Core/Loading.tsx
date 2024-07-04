import { Box, CircularProgress, Container, Typography } from "@mui/material";

export default function Loading() {
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
        <Typography component='h1' variant='h3' mb='16px'>
          Loading...
        </Typography>
        <CircularProgress />
      </Box>
    </Container>
  );
}
