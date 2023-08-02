package client

import (
	"context"
	"math/big"
	"pois_debug_client/pois_rpc"
	"time"

	"github.com/CESSProject/cess_pois/acc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func GetMinerPoisKey(minerId string) (acc.RsaKey, error) {
	conn, err := grpc.Dial("45.77.47.184:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return acc.RsaKey{}, err
	}
	defer conn.Close()
	cli := pois_rpc.NewPoisApiClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	response, err := cli.RequestMinerGetNewKey(ctx, &pois_rpc.RequestMinerInitParam{MinerId: []byte(minerId)})
	if err != nil {
		return acc.RsaKey{}, err
	}
	return acc.RsaKey{
		G: *big.NewInt(0).SetBytes(response.KeyG),
		N: *big.NewInt(0).SetBytes(response.KeyN),
	}, nil
}

func RegisterCommitConversational(minerId string) (bool, error) {
	conn, err := grpc.Dial("0.0.0.0:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return false, err
	}
	defer conn.Close()
	cli := pois_rpc.NewPoisApiClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	response, err := cli.RequestMinerRegister(ctx, &pois_rpc.RequestMinerInitParam{MinerId: []byte(minerId)})
	if err != nil {
		return false, err
	}
	return response.Status, nil
}

// func SendCommits(minerId string, commits []pois.Commit) ([][]int64, error) {
// 	conn, err := grpc.Dial("0.0.0.0:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer conn.Close()
// 	cli := pois_rpc.NewPoisApiClient(conn)
// 	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
// 	defer cancel()
// 	cms := make([]*pois_rpc.Commit, len(commits))
// 	for i := 0; i < len(commits); i++ {
// 		cms[i] = &pois_rpc.Commit{
// 			FileIndex: commits[i].FileIndex,
// 			Roots:     commits[i].Roots,
// 		}
// 	}
// 	response, err := cli.RequestMinerCommitGenChall(
// 		ctx, &pois_rpc.RequestMinerCommitGenChall{
// 			MinerId: []byte(minerId),
// 			Commit:  cms,
// 		},
// 	)
// 	chals := make([][]int64, len(commits))
// 	for i := 0; i < len(commits); i++ {
// 		chals[i] = response.Rows[i].Values
// 	}
// 	if err != nil {
// 		return nil, err
// 	}
// 	return chals, nil
// }

// func SendCommitProofs(minerId string, commitProofs [][]pois.CommitProof, accProofs *pois.AccProof) (bool, error) {
// 	maxMsgSize := 1024 * 1024 * 100
// 	conn, err := grpc.Dial(
// 		"0.0.0.0:50051",
// 		grpc.WithTransportCredentials(
// 			insecure.NewCredentials(),
// 		),
// 		grpc.WithDefaultCallOptions(
// 			grpc.MaxCallRecvMsgSize(maxMsgSize),
// 			grpc.MaxCallSendMsgSize(maxMsgSize),
// 		),
// 	)
// 	if err != nil {
// 		return false, err
// 	}
// 	defer conn.Close()
// 	cli := pois_rpc.NewPoisApiClient(conn)
// 	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
// 	defer cancel()
// 	commitGroup := &pois_rpc.CommitProofGroup{
// 		CommitProofGroupInner: make([]*pois_rpc.CommitProofGroupInner, len(commitProofs)),
// 	}
// 	for i := 0; i < len(commitProofs); i++ {
// 		commitGroup.CommitProofGroupInner[i] = &pois_rpc.CommitProofGroupInner{
// 			CommitProof: make([]*pois_rpc.CommitProof, len(commitProofs[i])),
// 		}
// 		for j := 0; j < len(commitProofs[i]); j++ {
// 			commitGroup.CommitProofGroupInner[i].CommitProof[j] = &pois_rpc.CommitProof{
// 				Node: &pois_rpc.MhtProof{
// 					Index: int32(commitProofs[i][j].Node.Index),
// 					Label: commitProofs[i][j].Node.Label,
// 					Paths: commitProofs[i][j].Node.Paths,
// 					Locs:  commitProofs[i][j].Node.Locs,
// 				},
// 				Parents: make([]*pois_rpc.MhtProof, len(commitProofs[i][j].Parents)),
// 			}
// 			for k := 0; k < len(commitProofs[i][j].Parents); k++ {
// 				commitGroup.CommitProofGroupInner[i].CommitProof[j].Parents[k] = &pois_rpc.MhtProof{
// 					Index: int32(commitProofs[i][j].Parents[k].Index),
// 					Label: commitProofs[i][j].Parents[k].Label,
// 					Paths: commitProofs[i][j].Parents[k].Paths,
// 					Locs:  commitProofs[i][j].Parents[k].Locs,
// 				}
// 			}
// 		}
// 	}

// 	accProof := &pois_rpc.AccProof{
// 		Indexs:  accProofs.Indexs,
// 		Labels:  accProofs.Labels,
// 		AccPath: accProofs.AccPath,
// 		WitChains: &pois_rpc.AccWitnessNode{
// 			Elem: accProofs.WitChains.Elem,
// 			Wit:  accProofs.WitChains.Wit,
// 			Acc: &pois_rpc.AccWitnessNode{
// 				Elem: accProofs.WitChains.Acc.Elem,
// 				Wit:  accProofs.WitChains.Acc.Wit,
// 				Acc: &pois_rpc.AccWitnessNode{
// 					Elem: accProofs.WitChains.Acc.Acc.Elem,
// 					Wit:  accProofs.WitChains.Acc.Acc.Wit,
// 				},
// 			},
// 		},
// 	}

// 	response, err := cli.RequestVerifyCommitProof(ctx,
// 		&pois_rpc.RequestVerifyCommitAndAccProof{
// 			MinerId:          []byte(minerId),
// 			CommitProofGroup: commitGroup,
// 			AccProof:         accProof,
// 		},
// 	)
// 	if err != nil {
// 		return false, err
// 	}
// 	return response.Status, nil
// }
