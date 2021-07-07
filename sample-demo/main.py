import sys
import grpc, yaml
import blockservice_pb2, blockservice_pb2_grpc, bsp_transaction_pb2
import proto.common.common_pb2 as common
from PySide6.QtWidgets import (QLineEdit, QPushButton, QApplication, QTextEdit,
    QVBoxLayout, QDialog,QLabel,QGridLayout,QComboBox,QSpinBox )
from  binascii import hexlify

config = {}
with open('config.yaml') as f:
    config = yaml.load(f, Loader=yaml.FullLoader)
    if config["bsp"]["endpoint"] == "":
        print("BSP endpoint must be given")
        sys.exit()
    print("Confirms that BSP endpoint is", config["bsp"]["endpoint"])


class Form(QDialog):
    def __init__(self, parent=None):
        super(Form, self).__init__(parent)
        # Create widgets

        self.edit = QLineEdit("1")
        self.region = QLineEdit("edgechain1")

        self.selectionBox = QComboBox()
        self.selectionBox.addItems(["Mainchain", "Edgechain"])
        self.selectionBox.currentTextChanged.connect(self.chain_change)

        self.regionBox = QComboBox()
        self.regionBox.addItems(["edgechain1", "edgechain2", "edgechain3"])
        self.regionBox.currentTextChanged.connect(self.region_change)

        self.button = QPushButton("Fetch Block")
        self.textEdit = QTextEdit()

        layout = QGridLayout()
        # layout.addWidget(self.selectionBox)
        # layout.addWidget(self.edit)
        # layout.addWidget(self.button)
        # layout.addWidget(self.textEdit)

        layout.addWidget(self.selectionBox, 0, 0)
        layout.addWidget(self.regionBox, 0, 1)

        layout.addWidget(self.edit, 1, 0, 1, 2)
        # layout.addWidget(self.region, 1, 1)

        layout.addWidget(self.button, 2, 0, 2, 2)
        # layout.addWidget(self.selectionBox, 0, 0)

        # Create layout and add widgets
        # layout = QVBoxLayout()

        # layout.addWidget(self.selectionBox)
        # layout.addWidget(self.edit)
        # layout.addWidget(self.button)
        layout.addWidget(self.textEdit, 4, 0, 6, 2)

        # Set dialog layout
        self.setLayout(layout)
        # self.setLayout(grid)
        # Add button signal to greetings slot
        self.button.clicked.connect(self.fetch_block)
        self.regionBox.setEnabled(False)
        self.client = MainchainClient(config["bsp"]["endpoint"])

    def chain_change(self, value):
        if value == "Mainchain":
            self.regionBox.setEnabled(False)
        else:
            self.regionBox.setEnabled(True)
        print("select", value)

    def region_change(self, value):
        print("select", value)

    def fetch_edgechain_block(self):
        print("fetch edgechain block!")

    def fetch_block(self):
        curr_chain = self.selectionBox.currentText()
        if curr_chain == "Edgechain":
            print("fetch from edgechain")
            self.fetch_edgechain_block()
        elif curr_chain == "Mainchain":
            print("fetch from mainchain")
            self.fetch_mainchain_block()

    def fetch_edgechain_block(self):
        region = self.regionBox.currentText()
        num_text = self.edit.text()
        if num_text.isdigit() == False:
            self.textEdit.setPlainText("Only positive integer is allowed")
            return
        resp = self.client.fetch_edgechain_block(int(num_text), region)
        hlf_block = common.Block()
        hlf_block.ParseFromString(resp.Block)
        print(resp)


    def fetch_mainchain_block(self):
        num_text = self.edit.text()
        if num_text.isdigit() == False:
            self.textEdit.setPlainText("Only positive integer is allowed")
            return

        resp = self.client.fetch_mainchain_block(int(num_text))
        mcheader = bsp_transaction_pb2.MCHeader()
        # print(resp)
        mcheader.ParseFromString(resp.Payload)

        mc_build_parameter = bsp_transaction_pb2.MCBuildParameter()
        mc_build_parameter.ParseFromString(mcheader.MCBuildParameter)

        print("ordering method", mc_build_parameter.OrderingMethod)

        method_params = bsp_transaction_pb2.BuildParameterOfIndexBasedOrdering()
        method_params.ParseFromString(mc_build_parameter.BuildParameter)

        resstr = "Round: {}\nNumber: {}\nPrevHash: {}\nMerkleRootHash: {}" \
                 "\nBucketSize: {}\n#TotalBlock: {}".format\
            (mcheader.Round, mcheader.Number, hexlify(bytearray(mcheader.PrevHash)), hexlify(bytearray(mcheader.MerkleRootHash)),
             method_params.BucketSize, len(mcheader.ECHeaders))

        self.textEdit.setPlainText(resstr)
        pecECBlocks = dict()

        for echeader in mcheader.ECHeaders:
            hlf_header = common.BlockHeader()
            hlf_header.ParseFromString(echeader.Header)
            pecECBlocks[echeader.Region] = list()

        for echeader in mcheader.ECHeaders:
            hlf_header = common.BlockHeader()
            hlf_header.ParseFromString(echeader.Header)
            pecECBlocks[echeader.Region].append(int(hlf_header.number))


        ectext = "\n***** EC-Header Membership Information ***** \n"
        for bsp in pecECBlocks:
            ectext += "\nBSP: {}, Blocks#: {}, from {} to {}".\
                format(bsp, len(pecECBlocks[bsp]), pecECBlocks[bsp][0], pecECBlocks[bsp][-1])

        resstr += "\nNumber of EC-Header {}".format(len(mcheader.ECHeaders))
        resstr += ectext

        self.textEdit.setPlainText(resstr)
        self.edit.setText(str(int(num_text)+1)) # Set next block


class MainchainClient:
    def __init__(self, endpoint, parent=None):
        self.endpoint = endpoint
        self.channel = grpc.insecure_channel(self.endpoint)
        self.BlockServiceStub = blockservice_pb2_grpc.BlockServiceStub(self.channel)
        print("Mainchain Client initialized to", self.endpoint)

    def fetch_mainchain_block(self, num):
        req = blockservice_pb2.BspBlockRequest(BlockNumber=num)
        return self.BlockServiceStub.GetMCBlockByNumber(req)

    def fetch_edgechain_block(self, num, region):
        req = blockservice_pb2.BspBlockRequest(BlockNumber=num, Region=region)
        return self.BlockServiceStub.GetECBlockByNumber(req)


if __name__ == '__main__':
    # Create the Qt Application
    app = QApplication(sys.argv)
    # Create and show the form
    form = Form()
    form.setWindowTitle('Mainchain Demo v0.1')

    form.show()
    # Run the main Qt loop
    sys.exit(app.exec())