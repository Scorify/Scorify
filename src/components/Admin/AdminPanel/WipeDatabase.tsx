import { useState } from "react";

import {
  Box,
  Button,
  Container,
  Paper,
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
  Typography,
  Tooltip,
} from "@mui/material";
import { enqueueSnackbar } from "notistack";

import { ConfirmModal } from "../..";
import { useWipeDatabaseMutation } from "../../../graph";

export default function WipeDatabase() {
  const [deleteUserCheckConfigurations, setDeleteUserCheckConfigurations] =
    useState(true);
  const [deleteInjectSubmissions, setDeleteInjectSubmissions] = useState(true);
  const [deleteStatusesScoresAndRounds, setDeleteStatusesScoresAndRounds] =
    useState(true);
  const [deleteCachedData, setDeleteCachedData] = useState(true);

  const [open, setOpen] = useState(false);
  const [wipeDatabase] = useWipeDatabaseMutation({
    onCompleted: () => {
      enqueueSnackbar("Database Wiped", { variant: "success" });
    },
    onError: (error) => {
      enqueueSnackbar(error.message, { variant: "error" });
    },
  });

  const databaseChanges = [
    {
      resource: "Users",
      action: "keep",
      color: "green",
    },
    {
      resource: "Inject Configurations",
      action: "keep",
      color: "green",
    },
    {
      resource: "Minion Configurations",
      action: "keep",
      color: "green",
    },
    {
      resource: "Admin Check Configurations",
      action: "keep",
      color: "green",
    },
    {
      resource: "User Check Configurations",
      action: deleteUserCheckConfigurations ? "delete" : "keep",
      color: deleteUserCheckConfigurations ? "red" : "green",
      toggle: () => setDeleteUserCheckConfigurations((prev) => !prev),
    },
    {
      resource: "User Inject Submissions",
      action: deleteInjectSubmissions ? "delete" : "keep",
      color: deleteInjectSubmissions ? "red" : "green",
      toggle: () => setDeleteInjectSubmissions((prev) => !prev),
    },
    {
      resource: "Score Check Statuses",
      action: deleteStatusesScoresAndRounds ? "delete" : "keep",
      color: deleteStatusesScoresAndRounds ? "red" : "green",
      toggle: () => setDeleteStatusesScoresAndRounds((prev) => !prev),
    },
    {
      resource: "Rounds",
      action: deleteStatusesScoresAndRounds ? "delete" : "keep",
      color: deleteStatusesScoresAndRounds ? "red" : "green",
      toggle: () => setDeleteStatusesScoresAndRounds((prev) => !prev),
    },
    {
      resource: "User Scores",
      action: deleteStatusesScoresAndRounds ? "delete" : "keep",
      color: deleteStatusesScoresAndRounds ? "red" : "green",
      toggle: () => setDeleteStatusesScoresAndRounds((prev) => !prev),
    },
    {
      resource: "All cached data",
      action: deleteCachedData ? "delete" : "keep",
      color: deleteCachedData ? "red" : "green",
      toggle: () => setDeleteCachedData((prev) => !prev),
    },
  ];

  const handleWipeDatabase = () => {
    wipeDatabase();
    setOpen(false);
  };

  return (
    <>
      <ConfirmModal
        title='Wipe Database Confirmation'
        subtitle={
          <Box>
            <Typography variant='h6' align='center'>
              This will WIPE ALL DATA from the database.
            </Typography>
            <Typography variant='h6' align='center'>
              This action CANNOT be undone.
            </Typography>
          </Box>
        }
        buttonText='Wipe Database'
        value='wipe database'
        label="Type 'wipe database' to confirm"
        onConfirm={handleWipeDatabase}
        open={open}
        setOpen={setOpen}
      />
      <Container maxWidth='sm'>
        <Typography variant='h4' align='center'>
          Wipe Database
        </Typography>
        <Box
          sx={{ m: 2 }}
          display='flex'
          alignItems='center'
          flexDirection='column'
        >
          <Typography variant='h6' textTransform='uppercase' fontWeight={900}>
            This will wipe all data from the database.
          </Typography>
          <Typography variant='h6' textTransform='uppercase' fontWeight={900}>
            This action cannot be undone.
          </Typography>
          <Box sx={{ m: 1 }} />
          <Typography variant='body1'>
            This will make the following changes:
          </Typography>

          <TableContainer
            component={Paper}
            sx={{ width: "fit-content", mt: 1, mb: 2 }}
          >
            <Table>
              <TableHead>
                <TableRow>
                  <TableCell size='small' align='center'>
                    <Typography variant='body1'>Action</Typography>
                  </TableCell>
                  <TableCell size='small' align='center'>
                    <Typography variant='body1'>Resource</Typography>
                  </TableCell>
                </TableRow>
              </TableHead>
              <TableBody>
                {databaseChanges.map((change, index) => (
                  <TableRow key={index}>
                    <TableCell size='small' align='center'>
                      {change.toggle ? (
                        <Tooltip
                          title={
                            <Box
                              display='flex'
                              flexDirection='column'
                              justifyContent='center'
                            >
                              <Typography variant='caption' align='center'>
                                Click to Toggle
                              </Typography>
                              <Typography variant='caption' align='center'>
                                This only edits what "wipe database" will do
                              </Typography>
                            </Box>
                          }
                        >
                          <Button
                            onClick={change.toggle}
                            variant='contained'
                            color={
                              change.color === "green" ? "success" : "error"
                            }
                          >
                            <Typography
                              variant='body2'
                              textTransform='uppercase'
                            >
                              {change.action}
                            </Typography>
                          </Button>
                        </Tooltip>
                      ) : (
                        <Typography
                          variant='body2'
                          style={{ color: change.color }}
                          textTransform='uppercase'
                        >
                          {change.action}
                        </Typography>
                      )}
                    </TableCell>
                    <TableCell size='small'>
                      <Typography variant='body2'>{change.resource}</Typography>
                    </TableCell>
                  </TableRow>
                ))}
              </TableBody>
            </Table>
          </TableContainer>
          <Button
            variant='contained'
            onClick={() => {
              setOpen(true);
            }}
          >
            <Typography variant='h6'>Wipe Database</Typography>
          </Button>
        </Box>
      </Container>
    </>
  );
}
