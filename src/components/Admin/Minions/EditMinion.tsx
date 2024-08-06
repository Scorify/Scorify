import { useEffect, useMemo, useState } from "react";

import {
  Box,
  Button,
  Chip,
  TextField,
  Tooltip,
  Typography,
} from "@mui/material";
import { Memory, Speed } from "@mui/icons-material";

import { Dropdown } from "../..";
import { MinionsQuery } from "../../../graph";

type props = {
  minion: MinionsQuery["minions"][0];
  handleRefetch: () => void;
  visible: boolean;
  sortMinions: () => void;
};

export default function EditCheck({ minion, visible, sortMinions }: props) {
  const [expanded, setExpanded] = useState(false);

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
        <Button variant='contained' color='error'>
          Deactivate
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

            // handleSave();
          }}
        >
          Save
        </Button>
      }
      toggleButtonVisible={nameChanged}
    >
      <Box
        sx={{
          display: "flex",
          gap: "16px",
          flexWrap: "wrap",
          justifyContent: "center",
        }}
      >
        {JSON.stringify(minion.metrics)}
      </Box>
    </Dropdown>
  );
}
