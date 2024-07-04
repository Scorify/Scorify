import { ReactElement, ReactNode, Suspense, useState } from "react";

import { ExpandLess, ExpandMore } from "@mui/icons-material";
import {
  Box,
  Card,
  CardContent,
  CardHeader,
  Collapse,
  Divider,
  Grow,
  IconButton,
  Slide,
} from "@mui/material";
import Loading from "../Core/Loading";

type props = {
  elevation?: number;
  modal?: ReactNode;
  title: ReactNode;
  expandableButtons?: ReactElement[];
  toggleButton?: ReactElement;
  toggleButtonVisible?: boolean;
  visible?: boolean;
  expanded: boolean;
  setExpanded: React.Dispatch<React.SetStateAction<boolean>>;
  children: ReactNode;
};

export default function Dropdown({
  elevation,
  modal,
  title,
  expandableButtons,
  toggleButton,
  toggleButtonVisible,
  visible = true,
  expanded,
  setExpanded,
  children,
}: props) {
  const [renderContent, setRenderContent] = useState(false);

  const toggleExpanded = () => setExpanded((prev) => !prev);

  return (
    <>
      {modal}
      <Grow in={true}>
        <Card
          sx={{
            width: "100%",
            marginBottom: "24px",
            display: visible ? "block" : "none",
          }}
          elevation={elevation}
        >
          <CardHeader
            title={
              <Box
                display='flex'
                flexDirection='row'
                alignItems='baseline'
                gap='12px'
              >
                {title}
              </Box>
            }
            action={
              <Box display='flex' flexDirection='row' gap='12px'>
                <Box
                  display='flex'
                  flexDirection='row'
                  gap='12px'
                  padding='0px 4px'
                  overflow='hidden'
                >
                  {expandableButtons &&
                    expandableButtons.map((button, i) => (
                      <Slide
                        key={i}
                        in={expanded}
                        direction='left'
                        unmountOnExit
                        mountOnEnter
                      >
                        {button}
                      </Slide>
                    ))}
                  {toggleButton && (
                    <Slide
                      in={toggleButtonVisible}
                      direction='left'
                      unmountOnExit
                      mountOnEnter
                    >
                      {toggleButton}
                    </Slide>
                  )}
                </Box>
                <IconButton>
                  {expanded ? <ExpandLess /> : <ExpandMore />}
                </IconButton>
              </Box>
            }
            onClick={toggleExpanded}
          />

          {expanded && <Divider sx={{ m: "0 1rem" }} />}

          <Collapse
            in={expanded}
            timeout={300}
            onEnter={() => {
              setRenderContent(true);
            }}
            onExited={() => {
              setRenderContent(false);
            }}
          >
            {renderContent && (
              <Suspense fallback={<Loading />}>
                <CardContent>{children}</CardContent>
              </Suspense>
            )}
          </Collapse>
        </Card>
      </Grow>
    </>
  );
}
