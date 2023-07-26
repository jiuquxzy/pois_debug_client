package test

import (
	"encoding/json"
	"math/big"
	"pois_debug_client/client"
	"testing"
	"time"

	"github.com/CESSProject/cess_pois/acc"
	"github.com/CESSProject/cess_pois/pois"
	"github.com/CESSProject/cess_pois/util"
)

var (
	minerId                   = "12D3KooWPBNzvjQtLEaDHBLACgZYBW9uHQyUtSvNoFnBm8GVPAJ4"
	k, n, d, space int64      = 7, 1024 * 16, 64, 1024 * 1024
	key            acc.RsaKey = acc.RsaKey{
		G: *big.NewInt(0).SetBytes([]byte{44, 68, 237, 11, 208, 232, 96, 106, 35, 14, 90, 202, 6, 79, 202, 102, 160, 4, 236, 113, 79, 255, 63, 125, 162, 182, 6, 253, 192, 72, 191, 14, 52, 19, 208, 190, 92, 219, 190, 93, 123, 17, 137, 243, 104, 119, 187, 6, 97, 244, 69, 167, 244, 196, 207, 165, 146, 178, 160, 234, 211, 237, 240, 103, 63, 108, 9, 223, 191, 70, 243, 13, 121, 90, 52, 146, 197, 77, 245, 84, 246, 202, 120, 57, 243, 245, 163, 1, 169, 61, 135, 198, 6, 78, 170, 88, 6, 235, 18, 222, 217, 16, 121, 61, 197, 114, 155, 47, 233, 34, 58, 62, 142, 140, 57, 48, 76, 123, 160, 38, 19, 200, 89, 64, 94, 250, 8, 187, 31, 151, 87, 12, 139, 28, 144, 123, 38, 41, 145, 153, 218, 7, 42, 238, 167, 208, 255, 189, 162, 58, 132, 180, 182, 232, 240, 158, 181, 222, 245, 213, 127, 226, 190, 250, 111, 54, 225, 140, 99, 131, 17, 10, 207, 102, 135, 57, 0, 26, 194, 118, 217, 183, 4, 82, 157, 67, 253, 150, 129, 155, 96, 152, 228, 181, 80, 76, 154, 92, 94, 157, 66, 99, 17, 40, 185, 110, 24, 147, 138, 204, 193, 148, 222, 55, 137, 149, 237, 247, 249, 2, 75, 179, 178, 231, 170, 147, 248, 189, 54, 78, 78, 196, 238, 170, 32, 237, 140, 35, 225, 36, 62, 210, 106, 19, 102, 253, 27, 69, 81, 68, 74, 205, 56, 1, 69, 225}),
		N: *big.NewInt(0).SetBytes([]byte{202, 251, 82, 59, 32, 166, 229, 179, 98, 47, 46, 251, 25, 171, 65, 194, 54, 140, 212, 243, 168, 12, 232, 40, 222, 183, 114, 188, 183, 248, 101, 9, 0, 255, 83, 151, 47, 32, 242, 83, 204, 98, 255, 179, 119, 45, 177, 119, 218, 195, 169, 51, 44, 198, 239, 235, 151, 147, 247, 93, 220, 137, 66, 53, 210, 187, 167, 15, 55, 6, 143, 59, 183, 107, 218, 125, 79, 46, 97, 200, 43, 61, 54, 208, 197, 189, 116, 82, 122, 178, 198, 81, 18, 108, 232, 226, 77, 180, 8, 76, 219, 28, 69, 9, 117, 145, 248, 127, 25, 141, 14, 173, 216, 54, 5, 139, 100, 76, 29, 88, 108, 123, 128, 88, 125, 96, 127, 254, 44, 39, 206, 66, 12, 184, 230, 82, 190, 107, 40, 113, 45, 10, 116, 77, 74, 45, 86, 162, 66, 34, 229, 200, 74, 62, 163, 55, 47, 69, 182, 181, 115, 155, 244, 87, 15, 138, 19, 189, 119, 150, 222, 215, 9, 180, 0, 217, 3, 254, 62, 170, 36, 241, 194, 122, 61, 122, 77, 238, 59, 160, 13, 88, 236, 98, 57, 59, 35, 222, 60, 177, 146, 140, 82, 187, 10, 240, 88, 148, 134, 146, 28, 35, 187, 115, 125, 228, 13, 0, 182, 216, 167, 234, 53, 222, 192, 66, 230, 202, 195, 115, 202, 12, 201, 58, 120, 49, 237, 90, 247, 69, 38, 222, 145, 118, 163, 164, 54, 222, 107, 49, 147, 126, 180, 63, 213, 103}),
	}
	chalsPath = "./commit_chals"
)

func TestRegisterConversational(t *testing.T) {

	resp, err := client.RegisterCommitConversational(minerId)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("register commit conversational success", resp)
}

func TestSendCommits(t *testing.T) {
	prover, err := pois.NewProver(k, n, d, []byte(minerId), space)
	if err != nil {
		t.Fatal(err)
	}
	err = prover.Init(key)
	if err != nil {
		t.Fatal(err)
	}
	prover.RunIdleFileGenerationServer(2)
	ok := prover.GenerateFile(16)
	if !ok {
		t.Fatal("generate file not ok")
	}
	time.Sleep(time.Second * 100)
	commits, err := prover.GetCommits(16)
	if err != nil {
		t.Fatal(err)
	}
	chals, err := client.SendCommits(minerId, commits)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("send commits success", chals)
	jbytes, err := json.Marshal(chals)
	if err != nil {
		t.Fatal(err)
	}
	err = util.SaveFile(chalsPath, jbytes)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSendCommitProofs(t *testing.T) {
	prover, err := pois.NewProver(k, n, d, []byte(minerId), space)
	if err != nil {
		t.Fatal(err)
	}
	err = prover.Recovery(key, 0, 0)
	if err != nil {
		t.Fatal(err)
	}
	prover.RunIdleFileGenerationServer(2)
	t.Log(prover.GetSpace())
	_, err = prover.GetCommits(16)
	if err != nil {
		t.Fatal(err)
	}
	data, err := util.ReadFile(chalsPath)
	if err != nil {
		t.Fatal(err)
	}
	var chals [][]int64
	if err := json.Unmarshal(data, &chals); err != nil {
		t.Fatal(err)
	}
	commitProof, accProof, err := prover.ProveCommitAndAcc(chals)
	if err != nil {
		t.Fatal(err)
	}
	ok, err := client.SendCommitProofs(minerId, commitProof, accProof)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("verify commit proofs and acc proof result", ok)
}
