import React, { forwardRef } from "react";
import ArrowUpward from "@mui/icons-material/ArrowUpward";
import ChevronLeft from "@mui/icons-material/ChevronLeft";
import ChevronRight from "@mui/icons-material/ChevronRight";
import Clear from "@mui/icons-material/Clear";
import FirstPage from "@mui/icons-material/FirstPage";
import LastPage from "@mui/icons-material/LastPage";
import Search from "@mui/icons-material/Search";
import MaterialTable from "material-table";

// n.b. this exists to bridge incompatibilities between the
// latest material-ui and material-table.
const icons = {
  DetailPanel: forwardRef((props, ref: React.RefObject<SVGSVGElement>) => (
    <ChevronRight {...props} ref={ref} />
  )),
  FirstPage: forwardRef((props, ref: React.RefObject<SVGSVGElement>) => (
    <FirstPage {...props} ref={ref} />
  )),
  LastPage: forwardRef((props, ref: React.RefObject<SVGSVGElement>) => (
    <LastPage {...props} ref={ref} />
  )),
  NextPage: forwardRef((props, ref: React.RefObject<SVGSVGElement>) => (
    <ChevronRight {...props} ref={ref} />
  )),
  PreviousPage: forwardRef((props, ref: React.RefObject<SVGSVGElement>) => (
    <ChevronLeft {...props} ref={ref} />
  )),
  ResetSearch: forwardRef((props, ref: React.RefObject<SVGSVGElement>) => (
    <Clear {...props} ref={ref} />
  )),
  Search: forwardRef((props, ref: React.RefObject<SVGSVGElement>) => (
    <Search {...props} ref={ref} />
  )),
  SortArrow: forwardRef((props, ref: React.RefObject<SVGSVGElement>) => (
    <ArrowUpward {...props} ref={ref} />
  )),
};

const flattenTreeData = (data: object, parentId?: string) => {
  let nodes = [];
  Object.keys(data).forEach(key => {
    let nodeId = key;
    if (parentId) {
      nodeId = `${parentId}-${key}`;
    }
    let value = data[key];
    if (value instanceof Object) {
      nodes = nodes.concat(flattenTreeData(value, nodeId));
      value = "";
    }
    const node = { id: nodeId, name: key, value, parentId: "" };
    if (parentId) {
      node.parentId = parentId;
    }
    nodes.push(node);
  });
  return nodes;
};

interface TreeTableProps {
  title?: string;
  data: object;
}

const TreeTable: React.FC<TreeTableProps> = ({ title, data }) => {
  const treeData = flattenTreeData(data);
  return (
    <MaterialTable
      title={title || ""}
      icons={icons}
      style={{ minWidth: "600px", width: "100%" }}
      options={{
        maxBodyHeight: "400px",
      }}
      data={treeData}
      columns={[
        { title: "Key", field: "name" },
        { title: "Value", field: "value" },
      ]}
      parentChildData={(row, rows) => rows.find(a => a.id === row.parentId)}
    />
  );
};

export default TreeTable;
