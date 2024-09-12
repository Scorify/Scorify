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
} from "@mui/material";
import { enqueueSnackbar } from "notistack";

import { useWipeDatabaseMutation } from "../../../graph";

function WipeDatabaseModal() {
  const [wipeDatabase] = useWipeDatabaseMutation({
    onCompleted: () => {
      enqueueSnackbar("Database Wiped", { variant: "success" });
    },
    onError: (error) => {
      enqueueSnackbar(error.message, { variant: "error" });
    },
  });

  return <></>;
}

export default function WipeDatabase() {
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
      action: "delete",
      color: "red",
    },
    {
      resource: "User Inject Submissions",
      action: "delete",
      color: "red",
    },
    {
      resource: "Score Check Statuses",
      action: "delete",
      color: "red",
    },
    {
      resource: "User Scores",
      action: "delete",
      color: "red",
    },
  ];

  return (
    <>
      <WipeDatabaseModal />
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
                      <Typography
                        variant='body2'
                        style={{ color: change.color }}
                        textTransform='uppercase'
                      >
                        {change.action}
                      </Typography>
                    </TableCell>
                    <TableCell size='small'>
                      <Typography variant='body2'>{change.resource}</Typography>
                    </TableCell>
                  </TableRow>
                ))}
              </TableBody>
            </Table>
          </TableContainer>
          <Button variant='contained' onClick={() => {}}>
            <Typography variant='h6'>Wipe Database</Typography>
          </Button>
        </Box>
      </Container>
    </>
  );
}
