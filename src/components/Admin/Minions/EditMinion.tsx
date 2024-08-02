import { useMemo, useState } from "react";

import { Box, Button, TextField, Typography } from "@mui/material";

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

  return (
    <Dropdown
      title={
        expanded ? (
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
        )
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
