import { useEffect, useMemo, useState } from "react";

import { Memory, Speed } from "@mui/icons-material";
import {
  Box,
  Button,
  Chip,
  TextField,
  Tooltip,
  Typography,
} from "@mui/material";
import { enqueueSnackbar } from "notistack";

import { Dropdown, Error, Loading, StatusTable } from "../..";
import {
  MinionsQuery,
  StatusEnum,
  StatusesQuery,
  useMinionStatusSummaryQuery,
  useStatusesQuery,
  useUpdateMinionMutation,
} from "../../../graph";

type props = {
  minion: MinionsQuery["minions"][0];
  handleRefetch: () => void;
  visible: boolean;
  sortMinions?: () => void;
  elevation?: number;
};

export default function EditCheck({
  minion,
  visible,
  sortMinions,
  handleRefetch,
  elevation,
}: props) {
  const [expanded, setExpanded] = useState(false);

  const [updateMinion] = useUpdateMinionMutation({
    onCompleted: () => {
      enqueueSnackbar("Minion updated successfully", { variant: "success" });
      handleRefetch();
    },
    onError: (error) => {
      enqueueSnackbar(error.message, { variant: "error" });
    },
  });

  const [name, setName] = useState<string>(minion.name);
  const nameChanged = useMemo(() => name !== minion.name, [name, minion.name]);

  const minionLastUpdated = new Date(minion.metrics?.timestamp).getTime();
  const [timeDifference, setTimeDifference] = useState<number>(
    Date.now() - minionLastUpdated
  );
  const [minionsSorted, setMinionsSorted] = useState(false);

  useEffect(() => {
    const interval = setInterval(() => {
      setTimeDifference(Date.now() - minionLastUpdated);
      if (
        sortMinions &&
        !minionsSorted &&
        Date.now() - minionLastUpdated > 60000
      ) {
        sortMinions();

        setMinionsSorted(true);
      }
    }, 1000);

    return () => clearInterval(interval);
    // eslint-disable-next-line react-hooks/exhaustive-deps -- sortMinions/minionsSorted read via closure, interval should only restart on timestamp change
  }, [minionLastUpdated]);

  const getMinionLastSeenLabel = (diff: number) => {
    if (!minion.metrics) {
      return "Never";
    }

    if (diff < 10000) {
      return "Just now";
    } else if (diff < 60000) {
      return `${Math.floor(diff / 1000)} seconds ago`;
    } else if (diff < 3600000) {
      return `${Math.floor(diff / 60000)} minutes ago`;
    } else if (diff < 86400000) {
      return `${Math.floor(diff / 3600000)} hours ago`;
    } else {
      return `${Math.floor(diff / 86400000)} days ago`;
    }
  };

  const bytesToSize = (bytes: number) => {
    const sizes = ["Bytes", "KB", "MB", "GB", "TB"];

    if (bytes === 0) {
      return "0 Byte";
    }

    const i = Math.floor(Math.log(bytes) / Math.log(1024));

    return `${parseFloat((bytes / Math.pow(1024, i)).toFixed(2))}${sizes[i]}`;
  };

  return (
    <Dropdown
      elevation={elevation}
      title={
        <>
          {expanded ? (
            <TextField
              label='Name'
              value={name}
              onClick={(e) => {
                e.stopPropagation();
              }}
              onChange={(e) => {
                setName(e.target.value);
              }}
              sx={{ marginRight: "12px" }}
              size='small'
            />
          ) : (
            <Typography variant='h6' component='div' marginRight='12px'>
              {minion.name}
            </Typography>
          )}
          <Box display='flex' alignItems='center' gap='8px'>
            <Tooltip title={`Last Seen: ${minionLastUpdated.toLocaleString()}`}>
              <Chip
                label={`${getMinionLastSeenLabel(timeDifference)}`}
                color={
                  Date.now() - minionLastUpdated < 60000 ? "success" : "error"
                }
                size='small'
              />
            </Tooltip>
            <Tooltip title={`IP Address: ${minion.ip}`}>
              <Chip label={minion.ip} size='small' />
            </Tooltip>
            {minion.metrics && (
              <>
                <Tooltip
                  title={`CPU Usage: ${minion.metrics.cpu_usage.toFixed(2)}%`}
                >
                  <Chip
                    icon={<Speed />}
                    label={`${minion.metrics.cpu_usage.toFixed(2)}%`}
                    size='small'
                    color={
                      minion.metrics.cpu_usage < 25
                        ? "success"
                        : minion.metrics.cpu_usage < 50
                        ? "warning"
                        : "error"
                    }
                  />
                </Tooltip>
                <Tooltip
                  title={`Memory Usage: ${bytesToSize(
                    minion.metrics.memory_usage
                  )} / ${bytesToSize(minion.metrics.memory_total)}`}
                >
                  <Chip
                    icon={<Memory />}
                    label={`${(
                      (minion.metrics.memory_usage /
                        minion.metrics.memory_total) *
                      100
                    ).toFixed(2)}%`}
                    size='small'
                    color={
                      minion.metrics.memory_usage /
                        minion.metrics.memory_total <
                      0.25
                        ? "success"
                        : minion.metrics.memory_usage /
                            minion.metrics.memory_total <
                          0.5
                        ? "warning"
                        : "error"
                    }
                  />
                </Tooltip>
              </>
            )}
          </Box>
        </>
      }
      expandableButtons={[
        <Button
          variant='contained'
          color={minion.deactivated ? "success" : "error"}
          onClick={(e) => {
            e.stopPropagation();

            updateMinion({
              variables: {
                id: minion.id,
                deactivated: !minion.deactivated,
              },
            });
          }}
        >
          {minion.deactivated ? "Activate" : "Deactivate"}
        </Button>,
      ]}
      visible={visible}
      expanded={expanded}
      setExpanded={setExpanded}
      toggleButton={
        <Button
          variant='contained'
          color='success'
          onClick={(e) => {
            if (!expanded) {
              e.stopPropagation();
            }

            updateMinion({
              variables: {
                id: minion.id,
                name: name,
              },
            });
          }}
        >
          Save
        </Button>
      }
      toggleButtonVisible={nameChanged}
    >
      <EditMinionChildren minion={minion} />
    </Dropdown>
  );
}

type editMinionChildrenProps = {
  minion: MinionsQuery["minions"][0];
};

function EditMinionChildren({ minion }: editMinionChildrenProps) {
  const [limit, setLimit] = useState<number>(10);
  const [statuses, setStatuses] = useState<StatusesQuery["statuses"]>([]);
  const { data, loading, error, refetch } = useStatusesQuery({
    variables: {
      statusesInputQuery: {
        minions: [minion.id],
        limit: limit,
      },
    },
  });

  const {
    data: summaryData,
    loading: summaryLoading,
    error: summaryError,
    refetch: summaryRefetch,
  } = useMinionStatusSummaryQuery({
    variables: {
      minion_id: minion.id,
    },
  });

  useEffect(() => {
    if (data) {
      setStatuses(data.statuses);
    }
  }, [data]);

  const handleLoadMore = () => {
    setLimit(limit + 10);
    refetch();
    summaryRefetch();
  };

  if (loading || summaryLoading) {
    return <Loading />;
  }

  if (error) {
    console.error(error);
    return <Error code={error.name} message={error.message} />;
  }

  if (summaryError) {
    console.error(summaryError);
    return <Error code={summaryError.name} message={summaryError.message} />;
  }

  return (
    <Box>
      <Typography variant='caption'>
        Showing {statuses.length}/
        {summaryData?.minionStatusSummary.total ?? "??"} statuses (
        <Typography color='lightgreen' variant='caption'>
          {statuses.filter((s) => s.status === StatusEnum.Up).length}/
          {summaryData?.minionStatusSummary.up ?? "??"} Up,{" "}
        </Typography>
        <Typography color='red' variant='caption'>
          {statuses.filter((s) => s.status === StatusEnum.Down).length}/
          {summaryData?.minionStatusSummary.down ?? "??"} Down,{" "}
        </Typography>
        <Typography color='orange' variant='caption'>
          {statuses.filter((s) => s.status === StatusEnum.Unknown).length}/
          {summaryData?.minionStatusSummary.unknown ?? "??"} Unknown
        </Typography>
        )
      </Typography>
      <StatusTable
        statuses={statuses}
        sx={{ position: "relative", maxHeight: "400px" }}
      />
      <Button
        variant='contained'
        onClick={handleLoadMore}
        sx={{ mt: "16px" }}
        disabled={
          statuses.length >= (summaryData?.minionStatusSummary.total || limit)
        }
        fullWidth
      >
        Load More Statuses
      </Button>
    </Box>
  );
}
