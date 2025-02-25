import { useContext, useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";

import {
  KeyboardArrowLeft,
  KeyboardDoubleArrowLeft,
} from "@mui/icons-material";
import {
  Box,
  Button,
  CircularProgress,
  Container,
  Typography,
} from "@mui/material";
import { enqueueSnackbar } from "notistack";

import { KothScoreboardWrapper } from "../../components";
import { AuthContext } from "../../components/Context";
import { NormalScoreboardTheme } from "../../constants";
import {
  KothScoreboardQuery,
  Role,
  useKothScoreboardQuery,
  useKothScoreboardUpdateSubscription,
} from "../../graph";

type props = {
  theme: "dark" | "light";
};

export default function KothScoreboardPage({ theme }: props) {
  const navigate = useNavigate();

  const { me } = useContext(AuthContext);

  const { data: rawData, error, loading, refetch } = useKothScoreboardQuery();
  const [data, setData] = useState<
    KothScoreboardQuery["kothScoreboard"] | undefined
  >(rawData?.kothScoreboard || undefined);

  useEffect(() => {
    refetch();
    refetch();
  }, []);

  useEffect(() => {
    if (rawData?.kothScoreboard) {
      setData(rawData.kothScoreboard);
    }
  }, [rawData]);

  useKothScoreboardUpdateSubscription({
    onData: (data) => {
      if (data.data.data?.kothScoreboardUpdate) {
        setData(data.data.data.kothScoreboardUpdate || undefined);
      }
    },
    onError: (error) => {
      console.error(error);
    },
  });

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
          {data?.round.number && data.round.number > 10 ? (
            <KeyboardDoubleArrowLeft
              sx={{ cursor: "pointer" }}
              onClick={() => {
                navigate(`/koth-scoreboard/${data?.round.number - 10}`);
              }}
            />
          ) : (
            <KeyboardDoubleArrowLeft sx={{ visibility: "hidden" }} />
          )}
          {data?.round.number && data.round.number > 1 ? (
            <KeyboardArrowLeft
              sx={{ cursor: "pointer" }}
              onClick={() => {
                navigate(`/koth-scoreboard/${data?.round.number - 1}`);
              }}
            />
          ) : (
            <KeyboardArrowLeft sx={{ visibility: "hidden" }} />
          )}
          <Box marginLeft={0.5} marginRight={0.5}>
            <Typography component='h1' variant='h5'>
              Round {data?.round.number}
            </Typography>
          </Box>
          <KeyboardArrowLeft sx={{ visibility: "hidden" }} />
          <KeyboardDoubleArrowLeft sx={{ visibility: "hidden" }} />
        </Box>
        {me && me.me && me.me.role === Role.User && (
          <Button
            variant='contained'
            color='primary'
            sx={{ marginTop: 2 }}
            size='small'
            onClick={(e) => {
              e.stopPropagation();
              navigator.clipboard.writeText(me.me?.id || "");
              enqueueSnackbar("User ID copied to clipboard", {
                variant: "success",
              });
            }}
          >
            Copy User ID
          </Button>
        )}
        <Box m={2} />
        {error && <Typography variant='h6'>Error: {error.message}</Typography>}
        {loading && !data && <CircularProgress />}
        {data && (
          <KothScoreboardWrapper
            theme={theme}
            data={data}
            scoreboardTheme={NormalScoreboardTheme}
          />
        )}
      </Box>
    </Container>
  );
}
