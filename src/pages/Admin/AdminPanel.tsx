import { useState } from "react";

import { Box, Button, ButtonGroup, Container, Typography } from "@mui/material";

import { useNavigate } from "react-router-dom";
import {
  EngineState as EngineStateComponent,
  Notification,
} from "../../components";
import { EngineState } from "../../graph";

type props = {
  engineState: EngineState | undefined;
};

export default function AdminPanel({ engineState }: props) {
  const navigate = useNavigate();
  const urlParams = new URLSearchParams(location.search);

  let param = urlParams.get("selected") || "notification";
  if (!["notification", "engine"].includes(param)) {
    param = "notification";
  }

  const [selected, setSelected] = useState<"notification" | "engine">(
    param as "notification" | "engine"
  );

  const components = {
    notification: <Notification />,
    engine: <EngineStateComponent engineState={engineState} />,
  };

  return (
    <Container component='main' maxWidth='xl'>
      <Container
        maxWidth='sm'
        sx={{
          mt: 8,
          mb: 2,
          display: "flex",
          flexDirection: "column",
          alignItems: "center",
        }}
      >
        <Typography component='h1' variant='h3'>
          Admin Panel
        </Typography>
        <Box sx={{ m: 2 }} />
        <ButtonGroup variant='contained' fullWidth sx={{ display: "flex" }}>
          <Button
            disabled={selected === "notification"}
            sx={{ flexGrow: 1 }}
            onClick={() => {
              setSelected("notification");

              urlParams.set("selected", "notification");
              navigate(`?${urlParams.toString()}`);
            }}
          >
            Send Notification
          </Button>
          <Button
            disabled={selected === "engine"}
            sx={{ flexGrow: 1 }}
            onClick={() => {
              setSelected("engine");

              urlParams.set("selected", "engine");
              navigate(`?${urlParams.toString()}`);
            }}
          >
            Engine State
          </Button>
        </ButtonGroup>
      </Container>
      {components[selected]}
    </Container>
  );
}
