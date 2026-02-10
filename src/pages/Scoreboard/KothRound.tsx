import { useEffect, useState } from "react";
import { useNavigate, useParams } from "react-router-dom";

import {
  KeyboardArrowLeft,
  KeyboardArrowRight,
  KeyboardDoubleArrowLeft,
  KeyboardDoubleArrowRight,
} from "@mui/icons-material";
import { Box, CircularProgress, Container, Typography } from "@mui/material";

import {
  LatestRoundSubscription,
  useKothScoreboardQuery,
  useLatestRoundSubscription,
} from "../../graph";
import { KothScoreboardWrapper } from "../../components";
import { NormalScoreboardTheme } from "../../constants";

type params = {
  round: string;
};

type props = {
  theme: "dark" | "light";
};

export default function KothScoreboardRoundPage({ theme }: props) {
  const { round } = useParams<params>();
  const navigate = useNavigate();

  const [latestRound, setLatestRound] =
    useState<LatestRoundSubscription["latestRound"]["number"]>();

  const { data, error, loading, refetch } = useKothScoreboardQuery({
    variables: { round: round ? parseInt(round) : undefined },
  });

  useEffect(() => {
    refetch();
    refetch();
    // eslint-disable-next-line react-hooks/exhaustive-deps -- intentional mount-only effect
  }, []);

  useLatestRoundSubscription({
    onError: (error) => {
      console.error(error);
    },
    onData: (data) => {
      if (
        data.data?.data?.latestRound.number &&
        round &&
        parseInt(round) === data.data?.data?.latestRound.number
      ) {
        navigate("/koth-scoreboard");
      }

      setLatestRound(data.data?.data?.latestRound.number);
    },
  });

  useEffect(() => {
    if (latestRound && round) {
      if (latestRound <= parseInt(round)) {
        navigate("/koth-scoreboard");
      }
    }
  }, [latestRound, round, navigate]);

  return (
    <Container component='main' maxWidth='xl'>
      <Box
        sx={{
          display: "flex",
          flexDirection: "column",
          alignItems: "center",
        }}
      >
        <Typography
          component='h1'
          variant='h3'
          fontWeight={700}
          sx={{
            marginTop: 5,
          }}
        >
          KoTH Scoreboard
        </Typography>
        <Box
          sx={{
            display: "flex",
            alignItems: "center",
          }}
        >
          {data?.kothScoreboard.round.number &&
          data.kothScoreboard.round.number > 10 ? (
            <KeyboardDoubleArrowLeft
              sx={{ cursor: "pointer" }}
              onClick={() => {
                navigate(
                  `/koth-scoreboard/${data?.kothScoreboard.round.number - 10}`
                );
              }}
            />
          ) : (
            <KeyboardDoubleArrowLeft sx={{ visibility: "hidden" }} />
          )}
          {data?.kothScoreboard.round.number &&
          data.kothScoreboard.round.number > 1 ? (
            <KeyboardArrowLeft
              sx={{ cursor: "pointer" }}
              onClick={() => {
                navigate(
                  `/koth-scoreboard/${data.kothScoreboard.round.number - 1}`
                );
              }}
            />
          ) : (
            <KeyboardArrowLeft sx={{ visibility: "hidden" }} />
          )}
          <Box marginLeft={0.5} marginRight={0.5}>
            <Typography
              component='h1'
              variant='h5'
              onClick={() => {
                navigate("/koth-scoreboard");
              }}
              sx={{ cursor: "pointer" }}
            >
              Round {data?.kothScoreboard.round.number}
            </Typography>
          </Box>
          {latestRound &&
          data?.kothScoreboard.round.number &&
          latestRound >= data?.kothScoreboard.round.number + 1 ? (
            <KeyboardArrowRight
              sx={{ cursor: "pointer" }}
              onClick={() => {
                navigate(
                  `/koth-scoreboard/${data?.kothScoreboard.round.number + 1}`
                );
              }}
            />
          ) : (
            <KeyboardArrowRight sx={{ visibility: "hidden" }} />
          )}
          {latestRound &&
          data?.kothScoreboard.round.number &&
          latestRound >= data?.kothScoreboard.round.number + 10 ? (
            <KeyboardDoubleArrowRight
              sx={{ cursor: "pointer" }}
              onClick={() => {
                navigate(
                  `/koth-scoreboard/${data?.kothScoreboard.round.number + 10}`
                );
              }}
            />
          ) : (
            <KeyboardDoubleArrowRight sx={{ visibility: "hidden" }} />
          )}
        </Box>
        <Box m={2} />
        {error && <Typography variant='h6'>Error: {error.message}</Typography>}
        {loading && !data && <CircularProgress />}
        {data && (
          <KothScoreboardWrapper
            theme={theme}
            data={data["kothScoreboard"]}
            scoreboardTheme={NormalScoreboardTheme}
          />
        )}
      </Box>
    </Container>
  );
}
