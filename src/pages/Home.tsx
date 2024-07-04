import { Avatar, Box, Container, Typography } from "@mui/material";

export default function Index() {
  return (
    <Container component='main' maxWidth='xs'>
      <Box
        sx={{
          marginTop: 8,
          display: "flex",
          flexDirection: "column",
          alignItems: "center",
        }}
      >
        <Avatar
          src={"/scorify.svg"}
          alt='logo'
          sx={{ width: "100%", height: "100%" }}
        />
        <Typography component='h1' variant='h3'>
          Welcome to the
        </Typography>
        <Typography component='h1' variant='h2' sx={{ fontWeight: "bold" }}>
          Scorify!
        </Typography>
      </Box>
    </Container>
  );
}
