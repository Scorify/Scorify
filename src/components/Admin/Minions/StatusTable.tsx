import {
  TableContainer,
  Table,
  TableHead,
  TableRow,
  TableCell,
  TableBody,
  Paper,
  Typography,
  SxProps,
} from "@mui/material";

import { GetStatusesMutation, StatusEnum } from "../../../graph";
import { NormalScoreboardTheme } from "../../../constants";

type props = {
  statuses: GetStatusesMutation["statuses"];
  sx?: SxProps;
};

export default function StatusTable({ statuses, sx }: props) {
  return (
    <TableContainer component={Paper} sx={sx}>
      <Table sx={{ width: "100%" }} stickyHeader>
        <TableHead>
          <TableRow>
            <TableCell size='small'>
              <Typography variant='body2' align='center'>
                Status
              </Typography>
            </TableCell>
            <TableCell size='small'>
              <Typography variant='body2' align='center'>
                Timestamp
              </Typography>
            </TableCell>
            <TableCell size='small'>
              <Typography variant='body2' align='center'>
                Team
              </Typography>
            </TableCell>
            <TableCell size='small'>
              <Typography variant='body2' align='center'>
                Check
              </Typography>
            </TableCell>
            <TableCell size='small'>
              <Typography variant='body2' align='center'>
                Round
              </Typography>
            </TableCell>
            <TableCell size='small'>
              <Typography variant='body2' align='center'>
                Error
              </Typography>
            </TableCell>
          </TableRow>
        </TableHead>
        <TableBody>
          {statuses.map((status) => (
            <TableRow key={status.id}>
              <TableCell
                size='small'
                sx={{
                  backgroundColor:
                    NormalScoreboardTheme.cell["dark"]["plain"][
                      status.status ?? StatusEnum.Unknown
                    ],
                }}
              ></TableCell>
              <TableCell size='small'>
                <Typography variant='body2' align='center'>
                  {new Date(status.update_time).toLocaleString()}
                </Typography>
              </TableCell>
              <TableCell size='small'>
                <Typography variant='body2' align='center'>
                  {status.user.username}
                </Typography>
              </TableCell>
              <TableCell size='small'>
                <Typography variant='body2' align='center'>
                  {status.check.name}
                </Typography>
              </TableCell>
              <TableCell size='small'>
                <Typography variant='body2' align='center'>
                  {status.round.number}
                </Typography>
              </TableCell>
              <TableCell size='small'>
                <Typography variant='body2' align='center'>
                  {status.error}
                </Typography>
              </TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>
    </TableContainer>
  );
}
