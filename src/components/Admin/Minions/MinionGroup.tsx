import { useState } from "react";

import { Typography, Box } from "@mui/material";

import { Dropdown, EditMinion } from "../..";
import { MinionsQuery } from "../../../graph";

type props = {
  title: string;
  minions: MinionsQuery["minions"];
  handleRefetch: () => void;
  sortMinions: () => void;
  search: string;
  emptyString: string;
  elevation?: number;
};

export default function MinionGroup({
  title,
  minions,
  search,
  handleRefetch,
  sortMinions,
  emptyString,
  elevation,
}: props) {
  const [expanded, setExpanded] = useState(false);
  return (
    <Dropdown
      title={title}
      expanded={expanded}
      setExpanded={setExpanded}
      elevation={elevation}
    >
      {minions &&
      minions.filter(
        (minion) =>
          minion.name &&
          minion.name.toLowerCase().includes(search.toLowerCase())
      ).length ? (
        minions.map((minion) => (
          <EditMinion
            key={minion.id}
            minion={minion}
            handleRefetch={handleRefetch}
            visible={minion.name.toLowerCase().includes(search.toLowerCase())}
            sortMinions={sortMinions}
            elevation={elevation === undefined ? undefined : elevation + 2}
          />
        ))
      ) : (
        <Typography component='h1' variant='h6' sx={{ mb: "16px" }}>
          {emptyString}
        </Typography>
      )}
      <Box mb='-24px' />
    </Dropdown>
  );
}
