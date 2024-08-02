import { useMemo, useState } from "react";

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
};

export default function EditCheck({ minion, visible }: props) {
  const [expanded, setExpanded] = useState(false);

  const [name, setName] = useState<string>(minion.name);
  const nameChanged = useMemo(() => name !== minion.name, [name, minion.name]);

  const minionLastUpdated = new Date(minion.metrics?.timestamp);
  const getMinionLastSeenLabel = () => {
    if (!minion.metrics) {
      return "Never";
    }

    const diff = Date.now() - minionLastUpdated.getTime();

    if (diff < 5000) {
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
              sx={{ marginRight: "24px" }}
              size='small'
            />
          ) : (
            <Typography variant='h6' component='div' marginRight='24px'>
              {minion.name}
            </Typography>
          )}
          <Tooltip title='Last Seen'>
            <Chip
              label={`${getMinionLastSeenLabel()}`}
              color={
                Date.now() - minionLastUpdated.getTime() < 60000
                  ? "success"
                  : "error"
              }
              size='small'
            />
          </Tooltip>
          <Tooltip title='IP Address'>
            <Chip label={minion.ip} size='small' />
          </Tooltip>
          {minion.metrics && (
            <>
              <Tooltip title='CPU Usage'>
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
              <Tooltip title='Memory Usage'>
                <Chip
                  icon={<Memory />}
                  label={`${bytesToSize(
                    minion.metrics.memory_usage
                  )} / ${bytesToSize(minion.metrics.memory_total)}`}
                  size='small'
                  color={
                    minion.metrics.memory_usage / minion.metrics.memory_total <
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
